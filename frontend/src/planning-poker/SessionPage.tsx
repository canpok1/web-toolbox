import { useCallback, useEffect, useState } from "react";
import { useParams, useSearchParams } from "react-router-dom";
import { useLoading } from "../common/hooks/useLoading";
import Alert from "./components/Alert";
import HostPanel, { type HostPanelEvent } from "./components/HostPanel";
import RoundSummaryPanel from "./components/RoundSummaryPanel";
import SessionInvitePanel from "./components/SessionInvitePanel";
import SessionSummaryPanel from "./components/SessionSummaryPanel";
import VotePanel from "./components/VotePanel";
import useSession from "./hooks/useSession";

function SessionPage() {
  const [searchParams] = useSearchParams();
  const { sessionId = "" } = useParams<{ sessionId: string }>();
  const participandId = searchParams.get("id") ?? "";
  const { session, round, myVote, myParticipant, loaded, error, reload } =
    useSession(sessionId, participandId);
  const [errorMessages, setErrorMessages] = useState<string[]>([]);
  const { setShowLoading } = useLoading();
  const [reconnectIntervalId, setReconnectIntervalId] =
    useState<NodeJS.Timeout | null>(null);

  const showHostPanel = session && participandId === session?.hostId;

  useEffect(() => {
    setShowLoading(!loaded);
  }, [setShowLoading, loaded]);

  const connectWebSocket = useCallback(() => {
    console.log("start websocket");
    try {
      // WebSocket 接続
      const ws = new WebSocket("/api/planning-poker/ws");
      console.log("made websocket");

      ws.onopen = () => {
        console.log("WebSocket connected");
      };

      ws.onmessage = async (event) => {
        console.log("Received message:", event.data);
        reload();
      };

      ws.onclose = () => {
        console.log("WebSocket disconnected");
        // 再接続を試みる
        const id = setInterval(() => {
          console.log("attempting to reconnect...");
          connectWebSocket();
        }, 3000);
        setReconnectIntervalId(id);
      };

      ws.onerror = (error) => {
        console.error("WebSocket error:", error);
      };

      // コンポーネントのアンマウント時に WebSocket 接続を閉じる
      return () => {
        console.log("call ws.close()");
        ws.close();
      };
    } catch (error) {
      console.error("websocket error: ", error);
      setErrorMessages(["エラーが発生しました。画面を再読み込みして下さい。"]);
    }
  }, [reload]);

  useEffect(() => {
    connectWebSocket();

    return () => {
      if (reconnectIntervalId) {
        clearInterval(reconnectIntervalId);
      }
    };
  }, [connectWebSocket, reconnectIntervalId]);

  useEffect(() => {
    if (error) {
      console.error(error);
      setErrorMessages(["エラーが発生しました。画面を再読み込みして下さい。"]);
      return;
    }

    if (!myParticipant) {
      setErrorMessages(["参加者が見つかりません。参加し直して下さい。"]);
      return;
    }

    setErrorMessages([]);
  }, [myParticipant, error]);

  return (
    <section className="mx-auto max-w-4xl px-5 py-5 text-center">
      <div className="mx-auto w-full">
        <h1 className="mb-5 font-bold text-3xl">プランニングポーカー</h1>

        <Alert className="w-full" messages={errorMessages} />

        {errorMessages.length === 0 && (
          <div className="mb-5 flex flex-col flex-wrap items-start justify-around gap-5 md:flex-row">
            {session && (
              <>
                <SessionSummaryPanel
                  className="w-full flex-2 md:max-w-lg"
                  session={session}
                  currentParticipantId={participandId}
                />
                <SessionInvitePanel
                  className="w-full flex-1 md:max-w-md"
                  sessionId={sessionId}
                />
              </>
            )}

            {showHostPanel && (
              <HostPanel
                className="w-full"
                session={session}
                round={round}
                onClick={(event: HostPanelEvent): void => {
                  console.log("clicked HostPanel, event:", event);
                  reload();
                }}
                onError={(event: HostPanelEvent, error: unknown): void => {
                  console.log(
                    "error on HostPanel, event:",
                    event,
                    "error:",
                    error,
                  );
                }}
              />
            )}

            {round?.status === "voting" && (
              <VotePanel
                className="w-full"
                roundId={round.roundId}
                participantId={participandId}
                scales={session?.scales ?? []}
                votedOption={myVote?.vote ?? null}
                onAfterVote={() => {
                  reload();
                }}
              />
            )}

            {round && (
              <RoundSummaryPanel
                className="w-full"
                participants={session?.participants ?? []}
                votes={round.votes}
                revealed={round?.status === "revealed"}
                summary={round.summary}
              />
            )}
          </div>
        )}
      </div>
    </section>
  );
}

export default SessionPage;
