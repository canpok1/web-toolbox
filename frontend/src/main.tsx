import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import CreateSessionPage from "./planning-poker/CreateSessionPage";
import JoinSessionPage from "./planning-poker/JoinSessionPage";
import PlanningPokerTopPage from "./planning-poker/PlanningPokerTopPage";
import SessionPage from "./planning-poker/SessionPage";
import "./index.css";
import TopPage from "./TopPage";

const rootElement = document.getElementById("root");

if (rootElement) {
  const root = createRoot(rootElement);
  root.render(
    <StrictMode>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<TopPage />} />
          <Route path="/planning-poker" element={<PlanningPokerTopPage />} />
          <Route
            path="/planning-poker/sessions/create"
            element={<CreateSessionPage />}
          />
          <Route
            path="/planning-poker/sessions/join"
            element={<JoinSessionPage />}
          />
          <Route
            path="/planning-poker/sessions/:sessionId"
            element={<SessionPage />}
          />
        </Routes>
      </BrowserRouter>
    </StrictMode>,
  );
} else {
  console.error("Could not find root element with id 'root'");
}
