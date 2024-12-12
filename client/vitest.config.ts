import { fileURLToPath } from 'node:url'
import { mergeConfig, defineConfig, configDefaults } from 'vitest/config'
import viteConfig from './vite.config'

export default mergeConfig(
  viteConfig,
  defineConfig({
    test: {
      environment: 'jsdom',
      exclude: [...configDefaults.exclude, 'e2e/**'],
      root: fileURLToPath(new URL('./', import.meta.url)),
      coverage: {
      },
      reporters: ['verbose', 'vitest-sonar-reporter'],
      outputFile: {
        'vitest-sonar-reporter': 'coverage/sonar-report.xml',
      }
    },
  }),
)
