import { CheckCircle2, Play, StopCircle } from "lucide-react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import RoundSummary from "./components/RoundSummary";
import SessionSummary from "./components/SessionSummary";
import VotePanel from "./components/VotePanel";
import useSession from "./hooks/useSession";
import type { RoundParticipant } from "./types/Participant";
import type { Round } from "./types/Round";

function SessionPage() {
  const { sessionId } = useParams<{ sessionId: string }>();
  const { session, fetch } = useSession();

  useEffect(() => {
    if (sessionId) {
      (async () => {
        await fetch(sessionId);
      })();
    }
  }, [fetch, sessionId]);

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
  const [round, setRound] = useState<Round | null>(null);
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

  const handleStartRound = () => {
    setRound({
      id: "xxxxx",
      status: "voting",
    });
  };

  const handleRevealVotes = () => {
    setRound({
      id: "xxxxx",
      status: "revealed",
    });
  };

  const handleVote = (option: string) => {
    setVoteOption(option);
    console.log(`voted: ${option}`);
  };

  return (
    <section className="mx-auto max-w-2xl px-5 py-25 text-center">
      <div className="mx-auto w-full">
        <h1 className="mb-5 font-bold text-3xl">プランニングポーカー</h1>
        {session && <SessionSummary session={session} />}

        <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
          <div className="card-body bg-neutral-content text-left">
            {round?.status !== "voting" && (
              <button
                type="button"
                className="btn btn-primary w-full"
                aria-label="ラウンドを開始"
                onClick={handleStartRound}
              >
                <Play />
                開始
              </button>
            )}
            {round?.status === "voting" && (
              <button
                type="button"
                className="btn btn-primary w-full"
                aria-label="投票を公開"
                onClick={handleRevealVotes}
              >
                <CheckCircle2 />
                投票を公開
              </button>
            )}
            {round?.status !== "voting" && (
              <button
                type="button"
                className="btn btn-error w-full"
                aria-label="セッションを終了"
              >
                <StopCircle />
                セッションを終了
              </button>
            )}
          </div>
        </div>

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
