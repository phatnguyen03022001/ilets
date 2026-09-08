#!/usr/bin/env python3
from __future__ import annotations

import argparse, hashlib, json, re, subprocess, sys
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parents[2]
SOURCE_MAP_PATH = ROOT / "tools/canonical/source-map.json"
SLICE_INPUT_DIR = ROOT / "tools/slice"
SLICE_OUTPUT_DIR = ROOT / "tools/slice/generated"
REGISTRY_OUTPUT = ROOT / "tools/canonical/generated/registry.json"
PROJECT_CATALOG_PATH = ROOT / "docs/catalog/project.json"
BEHAVIOR_PATH = ROOT / "docs/BEHAVIOR.md"

def fail(message: str) -> None: raise ValueError(message)
def git(*args: str) -> str: return subprocess.check_output(["git", *args], cwd=ROOT, text=True).strip()
def heading_level(heading: str) -> int:
    m = re.match(r"^(#+)\s", heading)
    if not m: fail(f"invalid heading in source map: {heading}")
    return len(m.group(1))
def section_text(text: str, heading: str, parent: str | None = None) -> str:
    lines=text.splitlines(); start_scope=0; end_scope=len(lines)
    if parent:
        pl=heading_level(parent)
        try: pi=lines.index(parent)
        except ValueError as exc: raise ValueError(f"missing parent section {parent!r}") from exc
        start_scope=pi+1
        for i in range(start_scope,len(lines)):
            if lines[i].startswith("#") and heading_level(lines[i]) <= pl: end_scope=i; break
    candidates=[i for i in range(start_scope,end_scope) if lines[i]==heading]
    if len(candidates)!=1: fail(f"expected exactly one section {heading!r} under {parent!r}; found {len(candidates)}")
    start=candidates[0]; level=heading_level(heading); end=end_scope
    for i in range(start+1,end_scope):
        if lines[i].startswith("#") and heading_level(lines[i]) <= level: end=i; break
    return "\n".join(lines[start:end])+"\n"
def strip_cell(value: str) -> str:
    value=value.strip()
    single_code_span=re.fullmatch(r"`([^`]*)`",value)
    if single_code_span:
        value=single_code_span.group(1)
    return value.strip()
def parse_table(section: str) -> dict[str,dict[str,Any]]:
    rows={}
    for raw in section.splitlines():
        line=raw.strip()
        if not (line.startswith("|") and line.endswith("|")): continue
        cells=[strip_cell(c) for c in line[1:-1].split("|")]
        if not cells or not cells[0] or set(cells[0]) <= {"-",":"}: continue
        candidate=cells[0]
        if candidate.lower() in {"id","node","condition","band"}: continue
        if candidate in rows: fail(f"duplicate canonical table id {candidate}")
        rows[candidate]={"id":candidate,"cells":cells[1:]}
    return rows
def parse_code_enum(section: str) -> dict[str,dict[str,Any]]:
    blocks=re.findall(r"```text\n(.*?)```",section,flags=re.DOTALL)
    if not blocks: fail("expected fenced text enum block")
    values={}
    for block in blocks:
        for line in block.splitlines():
            value=line.strip()
            if not value or not re.fullmatch(r"[A-Z][A-Z0-9_]*",value): continue
            if value in values: fail(f"duplicate enum value {value}")
            values[value]={"id":value,"cells":[]}
    return values
def source_blob(path: Path) -> str: return git("hash-object",str(path.relative_to(ROOT)))
REVISION_SCHEMA = "canonical-source-revision/v1"
def canonical_source_revision_from_inputs(source_map_bytes: bytes, owner_blobs: list[tuple[str,str]]) -> str:
    owners=[
        {"path":path,"git_blob":blob}
        for path,blob in sorted(set(owner_blobs))
    ]
    payload={
        "schema":REVISION_SCHEMA,
        "source_map_sha256":hashlib.sha256(source_map_bytes).hexdigest(),
        "owners":owners,
    }
    encoded=json.dumps(payload,sort_keys=True,separators=(",",":")).encode()
    return f"{REVISION_SCHEMA}:{hashlib.sha256(encoded).hexdigest()}"
def canonical_source_revision(owner_paths: list[str]) -> str:
    owner_blobs=[
        (owner,source_blob(ROOT/owner))
        for owner in sorted(set(owner_paths))
    ]
    return canonical_source_revision_from_inputs(SOURCE_MAP_PATH.read_bytes(),owner_blobs)
def slice_input_paths() -> list[Path]:
    return sorted(SLICE_INPUT_DIR.glob("*.json"), key=lambda path: path.as_posix())
def trace_output_path(slice_path: Path) -> Path:
    return SLICE_OUTPUT_DIR / f"{slice_path.stem}-trace.json"

def git_at(root: Path, *args: str) -> str:
    return subprocess.check_output(["git", *args], cwd=root, text=True).strip()

def revision_text(root: Path, revision: str, path: str) -> str:
    try:
        return subprocess.check_output(
            ["git", "show", f"{revision}:{path}"], cwd=root, text=True, stderr=subprocess.PIPE
        )
    except subprocess.CalledProcessError as exc:
        fail(f"missing referenced path: {path}")
        raise AssertionError from exc

def revision_json(root: Path, revision: str, path: str) -> Any:
    return json.loads(revision_text(root, revision, path))

def feature_alias_from_sources(feature_root_id: str, catalog: dict[str,Any], behavior: str) -> str:
    matches=[item for item in catalog.get("features",[]) if item.get("id")==feature_root_id]
    if len(matches)!=1: fail(f"expected exactly one canonical feature root {feature_root_id}; found {len(matches)}")
    feature=matches[0]
    if feature.get("spec_ref") != f"docs/BEHAVIOR.md#{feature_root_id}":
        fail(f"canonical feature root {feature_root_id} has conflicting behavior reference")
    name=feature.get("name")
    if not isinstance(name,str): fail(f"canonical feature root {feature_root_id} has malformed name")
    match=re.match(r"^([A-Z]-F\d{2})\s+—\s+",name)
    if not match: fail(f"canonical feature root {feature_root_id} has no lower-level feature alias")
    alias=match.group(1)
    headings=[line for line in behavior.splitlines() if re.match(rf"^#+\s+{re.escape(feature_root_id)}(?:\s|$)",line)]
    if len(headings)!=1: fail(f"expected exactly one behavior section for {feature_root_id}; found {len(headings)}")
    rows=parse_table(section_text(behavior,headings[0]))
    if alias not in rows: fail(f"behavior section for {feature_root_id} does not bind alias {alias}")
    return alias

def subject_slice_paths(root: Path, revision: str) -> list[str]:
    output=git_at(root,"ls-tree","-r","--name-only",revision,"--","tools/slice")
    paths=[]
    for path in output.splitlines():
        candidate=Path(path)
        if candidate.parent == Path("tools/slice") and candidate.suffix == ".json":
            paths.append(path)
    return sorted(paths)

def validated_relation_refs(items: Any, relation: str, read_subject: Any) -> list[dict[str,Any]]:
    if not isinstance(items,list) or not items: fail(f"{relation} must be a non-empty list")
    validated=[]
    for item in items:
        if not isinstance(item,dict) or set(item)!={"path","required_literals"}: fail(f"malformed {relation} relation")
        path=item["path"]; literals=item["required_literals"]
        if not isinstance(path,str) or not path or Path(path).is_absolute() or ".." in Path(path).parts: fail(f"malformed {relation} relation")
        if not isinstance(literals,list) or not literals or not all(isinstance(value,str) and value for value in literals): fail(f"malformed {relation} relation")
        text=read_subject(path)
        missing=[value for value in literals if value not in text]
        if missing: fail(f"unattributable {relation} relation {path}: missing {', '.join(missing)}")
        validated.append({"path":path,"required_literals":list(literals)})
    paths=[item["path"] for item in validated]
    if len(paths)!=len(set(paths)): fail(f"duplicate {relation} path")
    return sorted(validated,key=lambda item:item["path"])

def query_feature_trace(feature_root_id: str, expected_revision: str, root: Path = ROOT) -> dict[str,Any]:
    root=Path(root)
    subject=git_at(root,"rev-parse","HEAD")
    if not re.fullmatch(r"[0-9a-f]{40}",expected_revision): fail("expected revision must be a full 40-character Git SHA")
    if expected_revision != subject: fail(f"stale expected revision: expected {expected_revision}, subject {subject}")
    catalog=revision_json(root,subject,"docs/catalog/project.json")
    behavior=revision_text(root,subject,"docs/BEHAVIOR.md")
    alias=feature_alias_from_sources(feature_root_id,catalog,behavior)
    read_subject=lambda path: revision_text(root,subject,path)
    matches=[]
    for slice_path in subject_slice_paths(root,subject):
        slice_config=revision_json(root,subject,slice_path)
        if slice_config.get("feature_root_id")==feature_root_id:
            matches.append((slice_path,slice_config))
    common={
        "artifact":"DERIVED_DELIVERY_TRACE_QUERY",
        "schema_version":1,
        "authority":"NONE",
        "subject_revision":subject,
        "feature_root_id":feature_root_id,
        "aliases":[alias],
    }
    if not matches:
        return {
            **common,
            "condition":"IMPLEMENTATION_PROOF_NOT_ESTABLISHED",
            "implementation_refs":[],
            "verification_refs":[],
        }
    if len(matches)>1: fail(f"multiple slice traces for {feature_root_id}: {', '.join(path for path,_ in matches)}")
    _,slice_config=matches[0]
    if slice_config.get("feature_id")!=alias: fail(f"conflicting feature mapping for {feature_root_id}: expected {alias}, got {slice_config.get('feature_id')}")
    slice_id=slice_config.get("slice_id")
    if not isinstance(slice_id,str) or not slice_id: fail(f"malformed slice trace for {feature_root_id}: missing slice_id")
    implementation_refs=validated_relation_refs(slice_config.get("implementation_refs"),"implementation",read_subject)
    verification_refs=validated_relation_refs(slice_config.get("verification_refs"),"verification",read_subject)
    return {
        **common,
        "condition":"TRACE_PRESENT",
        "slice_id":slice_id,
        "implementation_refs":implementation_refs,
        "verification_refs":verification_refs,
    }
def registry_sha256(registry_doc: dict[str,Any]) -> str:
    return hashlib.sha256((json.dumps(registry_doc,sort_keys=True,separators=(",",":"))+"\n").encode()).hexdigest()
def validate_source_map(config: dict[str,Any]) -> None:
    allow=tuple(config["allowlisted_roots"]); forbidden=config["forbidden_authority_sources"]; seen=set()
    for registry in config["registries"]:
        owner=registry["owner"]
        if owner in forbidden or any(owner==item or owner.startswith(item.rstrip("/")+"/") for item in forbidden): fail(f"forbidden authority source: {owner}")
        if not any(owner==root or owner.startswith(root+"/") for root in allow): fail(f"owner outside canonical allowlist: {owner}")
        for object_id in registry["ids"]:
            pair=(registry["type"],object_id)
            if pair in seen: fail(f"duplicate configured canonical id {registry['type']}:{object_id}")
            seen.add(pair)
def materialize() -> tuple[dict[str,Any],list[tuple[Path,dict[str,Any]]]]:
    config=json.loads(SOURCE_MAP_PATH.read_text()); validate_source_map(config)
    owner_paths=sorted({e["owner"] for e in config["registries"]})
    revision=canonical_source_revision(owner_paths)
    entries=[]; global_ids={}
    for registry in config["registries"]:
        owner_path=ROOT/registry["owner"]
        if not owner_path.is_file(): fail(f"missing canonical owner: {registry['owner']}")
        section=section_text(owner_path.read_text(),registry["section"],registry.get("parent_section"))
        parsed=parse_table(section) if registry["kind"]=="markdown_table" else parse_code_enum(section) if registry["kind"]=="code_enum" else None
        if parsed is None: fail(f"unsupported source kind: {registry['kind']}")
        blob=source_blob(owner_path)
        for object_id in registry["ids"]:
            if object_id not in parsed: fail(f"unknown selected canonical reference {object_id} in {registry['owner']}::{registry['section']}")
            prior=global_ids.get(object_id)
            if prior and prior != registry["type"]: fail(f"canonical id {object_id} ambiguously materialized as {prior} and {registry['type']}")
            global_ids[object_id]=registry["type"]
            entries.append({"type":registry["type"],"id":object_id,"cells":parsed[object_id]["cells"],"source":{"owner":registry["owner"],"section":registry["section"],"parent_section":registry.get("parent_section"),"git_blob":blob}})
    entries.sort(key=lambda item:(item["type"],item["id"]))
    registry_doc={"artifact":"DERIVED_CANONICAL_REGISTRY","schema_version":1,"canonical_source_revision":revision,"source_map_sha256":hashlib.sha256(SOURCE_MAP_PATH.read_bytes()).hexdigest(),"entries":entries}
    traces=[]
    for slice_path in slice_input_paths():
        slice_config=json.loads(slice_path.read_text())
        refs={slice_config["feature_id"],slice_config["practice_mode_id"],*slice_config["practice_type_ids"],*slice_config["skill_target_ids"],*slice_config["official_family_ids"],slice_config["content_context_id"],slice_config["primary_activity_purpose"],slice_config["evidence_candidacy"]}
        missing=sorted(refs-set(global_ids))
        if missing: fail(f"slice trace {slice_path.name} contains unknown canonical references: {', '.join(missing)}")
        trace_doc={"artifact":"DERIVED_IMPLEMENTATION_TRACE","schema_version":1,"canonical_source_revision":revision,"canonical_registry_sha256":registry_sha256(registry_doc),"slice":slice_config}
        traces.append((slice_path,trace_doc))
    return registry_doc,traces
def render(doc: dict[str,Any]) -> str: return json.dumps(doc,indent=2,sort_keys=True)+"\n"
def write_or_check(path: Path, content: str, check: bool) -> None:
    if check:
        if not path.is_file(): fail(f"missing generated artifact: {path.relative_to(ROOT)}")
        if path.read_text()!=content: fail(f"generated artifact drift: {path.relative_to(ROOT)}")
    else:
        path.parent.mkdir(parents=True,exist_ok=True); path.write_text(content)
def main(argv: list[str] | None = None, root: Path = ROOT) -> int:
    parser=argparse.ArgumentParser()
    parser.add_argument("--check",action="store_true")
    parser.add_argument("--query-feature")
    parser.add_argument("--expected-revision")
    args=parser.parse_args(argv)
    try:
        if args.query_feature or args.expected_revision:
            if args.check or not args.query_feature or not args.expected_revision:
                fail("feature query requires --query-feature and --expected-revision without --check")
            print(render(query_feature_trace(args.query_feature,args.expected_revision,root=root)),end="")
            return 0
        registry_doc,traces=materialize(); write_or_check(REGISTRY_OUTPUT,render(registry_doc),args.check)
        for slice_path,trace_doc in traces: write_or_check(trace_output_path(slice_path),render(trace_doc),args.check)
    except (ValueError,subprocess.CalledProcessError,OSError,json.JSONDecodeError) as exc:
        print(f"canonical materialization failed: {exc}",file=sys.stderr); return 1
    return 0
if __name__=="__main__": raise SystemExit(main())
