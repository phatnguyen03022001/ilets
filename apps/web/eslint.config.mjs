import { defineConfig } from "eslint/config";
import nextVitals from "eslint-config-next/core-web-vitals";

function restoreEslint9ContextAccessors(rule) {
  if (!rule || typeof rule.create !== "function") {
    return rule;
  }

  return {
    ...rule,
    create(context) {
      if ("getFilename" in context) {
        return rule.create(context);
      }

      const legacyContext = Object.assign(Object.create(context), {
        getCwd: () => context.cwd,
        getFilename: () => context.filename,
        getPhysicalFilename: () => context.physicalFilename,
        getSourceCode: () => context.sourceCode,
        parserOptions: context.languageOptions.parserOptions,
      });

      return rule.create(Object.freeze(legacyContext));
    },
  };
}

const typescriptParser = nextVitals.find(
  (config) => config.name === "next/typescript",
)?.languageOptions?.parser;

if (!typescriptParser) {
  throw new Error("eslint-config-next did not expose its TypeScript parser");
}

function makeEslint10Compatible(configs) {
  const plugins = new WeakMap();

  return configs.map((config) => {
    const compatibleConfig =
      config.name === "next"
        ? {
            ...config,
            languageOptions: {
              ...config.languageOptions,
              parser: typescriptParser,
              parserOptions: { sourceType: "module" },
            },
          }
        : config;

    if (!compatibleConfig.plugins) {
      return compatibleConfig;
    }

    return {
      ...compatibleConfig,
      plugins: Object.fromEntries(
        Object.entries(compatibleConfig.plugins).map(([name, plugin]) => {
          if (!plugins.has(plugin)) {
            plugins.set(plugin, {
              ...plugin,
              rules: Object.fromEntries(
                Object.entries(plugin.rules ?? {}).map(([ruleName, rule]) => [
                  ruleName,
                  restoreEslint9ContextAccessors(rule),
                ]),
              ),
            });
          }

          return [name, plugins.get(plugin)];
        }),
      ),
    };
  });
}

// eslint-config-next 16.3.2 still combines plugins that call RuleContext
// accessors removed by ESLint 10 with Next's bundled Babel parser, whose scope
// manager lacks ESLint 10's addGlobals() API. Keep the selected Next rules and
// already-selected TypeScript parser, but adapt those legacy API boundaries
// until the upstream Next stack is natively ESLint 10 compatible.
export default defineConfig([...makeEslint10Compatible(nextVitals)]);
