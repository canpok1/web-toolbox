import { useEffect, useRef } from "react";
import { useParams, useSearchParams } from "react-router-dom";
import Alert from "./components/Alert";
import FibonacciVotePanel from "./components/FibonacciVotePanel";
import HostPanel, { type HostPanelEvent } from "./components/HostPanel";
import RoundSummary from "./components/RoundSummary";
import SessionSummary from "./components/SessionSummary";
import useSession from "./hooks/useSession";
import { ExtractErrorMessage } from "./utils/error";

function SessionPage() {
  const [searchParams] = useSearchParams();

  const { sessionId = "" } = useParams<{ sessionId: string }>();
  const participandId = searchParams.get("id") ?? "";

  const { session, round, myVote, error, reload } = useSession(
    sessionId,
    participandId,
  );
  const errorMessages = error ? [ExtractErrorMessage(error)] : [];

  const showHostPanel = session && participandId === session?.hostId;

  const intervalIdRef = useRef<NodeJS.Timeout | null>(null); // setIntervalのIDを保持

  // 定期更新
  useEffect(() => {
    intervalIdRef.current = setInterval(reload, 5000);
    return () => {
      if (intervalIdRef.current) {
        console.log(
          "useEffect cleanup: インターバルをクリアします",
          intervalIdRef.current,
        );
        clearInterval(intervalIdRef.current);
        intervalIdRef.current = null;
      }
    };
  }, [reload]);

  return (
    <section className="mx-auto max-w-2xl px-5 py-25 text-center">
      <div className="mx-auto w-full">
        <h1 className="mb-5 font-bold text-3xl">プランニングポーカー</h1>
        {session && (
          <SessionSummary
            session={session}
            currentParticipantId={participandId}
          />
        )}

        <Alert messages={errorMessages} />

        {showHostPanel && (
          <HostPanel
            session={session}
            round={round}
            onClick={(event: HostPanelEvent): void => {
              console.log("clicked HostPanel, event:", event);
              reload();
            }}
            onError={(event: HostPanelEvent, error: unknown): void => {
              console.log("error on HostPanel, event:", event, "error:", error);
            }}
          />
        )}

        {round?.status === "voting" && (
          <FibonacciVotePanel
            roundId={round.roundId}
            participantId={participandId}
            votedOption={myVote?.vote ?? null}
            onAfterVote={() => {
              reload();
            }}
          />
        )}

        {round && (
          <RoundSummary
            participants={session?.participants ?? []}
            votes={round.votes}
            revealed={round?.status === "revealed"}
          />
        )}
      </div>
    </section>
  );
}

export default SessionPage;
