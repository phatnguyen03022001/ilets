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

function makeEslint10Compatible(configs) {
  const plugins = new WeakMap();

  return configs.map((config) => {
    if (!config.plugins) {
      return config;
    }

    return {
      ...config,
      plugins: Object.fromEntries(
        Object.entries(config.plugins).map(([name, plugin]) => {
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

// eslint-config-next 16.3.2 still brings plugins that advertise ESLint <=9
// support and call RuleContext accessors removed by ESLint 10. Keep the Next
// rules enabled, but adapt that legacy API boundary until those plugins ship
// native ESLint 10 support.
export default defineConfig([...makeEslint10Compatible(nextVitals)]);
