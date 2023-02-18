/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_URL: string // backend url
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
