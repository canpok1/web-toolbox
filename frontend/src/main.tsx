import "./index.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Layout from "./Layout";
import TopPage from "./TopPage";
import CreateSessionPage from "./planning-poker/CreateSessionPage";
import JoinSessionPage from "./planning-poker/JoinSessionPage";
import PlanningPokerTopPage from "./planning-poker/PlanningPokerTopPage";
import SessionPage from "./planning-poker/SessionPage";

const rootElement = document.getElementById("root");

if (rootElement) {
  const root = createRoot(rootElement);
  root.render(
    <StrictMode>
      <BrowserRouter>
        <Routes>
          <Route element={<Layout />}>
            <Route
              path="/"
              element={
                <>
                  <title>Web Toolbox</title>
                  <TopPage />
                </>
              }
            />
            <Route
              path="/planning-poker"
              element={
                <>
                  <title>プランニングポーカー | Web Toolbox</title>
                  <PlanningPokerTopPage />
                </>
              }
            />
            <Route
              path="/planning-poker/sessions/create"
              element={
                <>
                  <title>プランニングポーカー | Web Toolbox</title>
                  <CreateSessionPage />
                </>
              }
            />
            <Route
              path="/planning-poker/sessions/join"
              element={
                <>
                  <title>プランニングポーカー | Web Toolbox</title>
                  <JoinSessionPage />
                </>
              }
            />
            <Route
              path="/planning-poker/sessions/:sessionId"
              element={
                <>
                  <title>プランニングポーカー | Web Toolbox</title>
                  <SessionPage />
                </>
              }
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </StrictMode>,
  );
} else {
  console.error("Could not find root element with id 'root'");
}
