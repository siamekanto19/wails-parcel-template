import { runtime } from "./wailsjs/runtime/runtime";

export {};
declare global {
  interface Window {
    runtime: runtime;
  }
}
