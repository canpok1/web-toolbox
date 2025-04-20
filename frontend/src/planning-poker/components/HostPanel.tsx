import { CheckCircle2, Play } from "lucide-react";
import { ApiClient } from "../../api/ApiClient";
import type { Round } from "../types/Round";
import type { Session } from "../types/Session";

export type HostPanelEvent = "startRound" | "revealVotes" | "endSession";

export type HostPanelProps = {
  session: Session | null;
  round: Round | null;
  onClick: (event: HostPanelEvent) => void;
  onError: (event: HostPanelEvent, error: unknown) => void;
  className?: string;
};

export default function HostPanel({
  session,
  round,
  onClick,
  onError,
  className,
}: HostPanelProps) {
  const showStartRoundButton = !round || round.status === "revealed";
  const showRevealVoteButton = round?.status === "voting";

  const handleStartRound = async () => {
    try {
      if (!session) {
        return;
      }
      const client = new ApiClient();
      await client.startRound(session.sessionId);
      onClick("startRound");
    } catch (error) {
      onError("startRound", error);
    }
  };

  const handleRevealVotes = async () => {
    try {
      if (!round) {
        return;
      }
      const client = new ApiClient();
      await client.revealRound(round.roundId);
      onClick("revealVotes");
    } catch (error) {
      onError("revealVotes", error);
    }
  };

  return (
    <div className={`card mx-auto shadow-sm ${className}`}>
      <div className="card-body bg-neutral-content text-left">
        {showStartRoundButton && (
          <button
            type="button"
            className="btn btn-primary w-full"
            aria-label="投票を開始"
            onClick={handleStartRound}
          >
            <Play />
            投票を開始
          </button>
        )}
        {showRevealVoteButton && (
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
      </div>
    </div>
  );
}
