module.exports = {
  env: {
    browser: true,
    es6: true,
    jest: true
    // es2021: true,
  },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:jest/recommended',
    'plugin:jsonc/recommended-with-jsonc'
  ],
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
  },
  overrides: [
    {
      files: ['*.ts','*.tsx'],
      parser: '@typescript-eslint/parser',
      parserOptions: {
        project: './tsconfig.json',
      }
    },
    {
      files: ['*.json'],
      rules: {
        'comma-dangle': ['error', 'never'],
        'jsonc/sort-keys': ['error',
          // For example, a definition for package.json
          {
            // Hits the root properties
            'order': [
              'name',
              'version',
              'private',
              'publishConfig'
              // ...
            ], 
            'pathPattern': '^$'
          },
          {
            'order': { 'type': 'asc' },
            'pathPattern': '^(?:dev|peer|optional|bundled)?[Dd]ependencies$'
          }
          // ...
        ],
        quotes: ['error', 'double']
      }
    },
    {
      env: {
        node:true
      },
      files: ['*.eslintrc.js'],
    },
  ],
  parserOptions: {
    ecmaFeatures: {
      jsx: true
    },
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  plugins: ['react', '@typescript-eslint','typescript-sort-keys','sort-keys-fix','jest'],
  rules: {
    /*"curly": ["error", "all"],
    "import/no-named-as-default": "off",
    "no-alert": "error",
    "no-console": "error",
    "no-redeclare": "off",
    "no-template-curly-in-string": "error",
    "no-var": "error",
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/interface-name-prefix": "off",
    "react/react-in-jsx-scope": "off",
    "@typescript-eslint/camelcase": "off",
    "@typescript-eslint/ban-ts-comment": "off",
    "@typescript-eslint/ban-types": [
      "error",
      {
        "extendDefaults": true,
        "types": {
          "{}": false
        }
      }
    ],
    "prefer-arrow-callback": "error",
    "@typescript-eslint/naming-convention": [
      "warn",
      {
        //Allow _ variable name
        "filter": {
          "match": false,
          "regex": "^_$"
        },

        "format": ["camelCase", "PascalCase", "UPPER_CASE"],

        "selector": "variable"
      }
    ],
    "prefer-const": "error",
    "@typescript-eslint/no-empty-function": "off",
    "prefer-destructuring": "warn",
    "@typescript-eslint/no-empty-interface": "warn",
    */
    '@typescript-eslint/no-unused-vars': 'warn',
    indent: ['warn', 2],
    'no-undef': ['error',],
    
    'no-unused-vars': 'warn',
    /*'linebreak-style': [
            'error',
            'unix'
        ],*/
    quotes: ['warn', 'single'],
    'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off',
    semi: ['warn', 'never'],
    'sort-keys-fix/sort-keys-fix':['warn', 'asc', {'natural': false}],
    'typescript-sort-keys/interface': 'warn',
    'typescript-sort-keys/string-enum': 'warn',
    '@typescript-eslint/no-empty-function': 'warn',
    '@typescript-eslint/ban-ts-comment': 'warn',
    'no-debugger': 'warn'
  }
}
