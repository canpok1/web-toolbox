import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

const rootElement = document.getElementById("root");

if (rootElement) {
  const root = createRoot(rootElement);
  root.render(
    <StrictMode>
      <h1>トップページ</h1>
    </StrictMode>,
  );
} else {
  console.error("Could not find root element with id 'root'");
}
