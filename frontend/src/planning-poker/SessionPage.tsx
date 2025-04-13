import { useEffect, useState } from "react";
import { useParams, useSearchParams } from "react-router-dom";
import HostPanel, { type HostPanelEvent } from "./components/HostPanel";
import RoundSummary from "./components/RoundSummary";
import SessionSummary from "./components/SessionSummary";
import VotePanel from "./components/VotePanel";
import useRound from "./hooks/useRound";
import useSession from "./hooks/useSession";
import type { RoundParticipant } from "./types/Participant";

function SessionPage() {
  const { sessionId } = useParams<{ sessionId: string }>();
  const { session, fetch: fetchSession } = useSession();
  const { round, fetch: fetchRound } = useRound();

  const [searchParams] = useSearchParams();
  const participandId = searchParams.get("id") ?? "";
  const showHostPanel = session && participandId === session?.hostId;

  useEffect(() => {
    if (sessionId) {
      (async () => {
        await fetchSession(sessionId);
      })();
    }
  }, [fetchSession, sessionId]);

  useEffect(() => {
    const roundId = session?.currentRoundId;
    if (roundId) {
      (async () => {
        await fetchRound(roundId);
      })();
    }
  }, [session, fetchRound]);

  const participants: RoundParticipant[] = [
    { id: "aaaa", name: "Aさん", vote: 1 },
    { id: "bbbb", name: "Bさん", vote: 2 },
    { id: "cccc", name: "Cさん", vote: null },
    { id: "dddd", name: "Dさん", vote: null },
    { id: "eeee", name: "Eさん", vote: 3 },
    { id: "ffff", name: "Fさん", vote: 1 },
    { id: "gggg", name: "Gさん", vote: 1 },
    { id: "hhhh", name: "Hさん", vote: 2 },
  ];
  const voteOptions = [
    "0",
    "1",
    "2",
    "3",
    "5",
    "8",
    "13",
    "21",
    "34",
    "55",
    "89",
    "?",
  ];
  const [voteOption, setVoteOption] = useState<string | null>(null);

  const handleVote = (option: string) => {
    setVoteOption(option);
    console.log(`voted: ${option}`);
  };

  return (
    <section className="mx-auto max-w-2xl px-5 py-25 text-center">
      <div className="mx-auto w-full">
        <h1 className="mb-5 font-bold text-3xl">プランニングポーカー</h1>
        {session && <SessionSummary session={session} />}

        {showHostPanel && (
          <HostPanel
            session={session}
            round={round}
            onClick={(event: HostPanelEvent): void => {
              console.log("clicked HostPanel, event:", event);
            }}
            onError={(event: HostPanelEvent, error: unknown): void => {
              console.log("error on HostPanel, event:", event, "error:", error);
            }}
          />
        )}

        {round?.status === "voting" && (
          <VotePanel
            voteOptions={voteOptions}
            onClick={handleVote}
            votedOption={voteOption}
          />
        )}

        {round && (
          <RoundSummary
            participants={participants}
            revealed={round?.status === "revealed"}
          />
        )}
      </div>
    </section>
  );
}

export default SessionPage;
