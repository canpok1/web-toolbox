import { CheckCircle2, Play, StopCircle } from "lucide-react";
import { ApiClient } from "../../api/ApiClient";
import type { Round } from "../types/Round";
import type { Session } from "../types/Session";

export type HostPanelEvent = "startRound" | "revealVotes" | "endSession";

export type HostPanelProps = {
  session: Session | null;
  round: Round | null;
  onClick: (event: HostPanelEvent) => void;
  onError: (event: HostPanelEvent, error: unknown) => void;
};

export default function HostPanel({
  session,
  round,
  onClick,
  onError,
}: HostPanelProps) {
  const hasEnableSession = session && session.status !== "finished";
  const showStartRoundButton = !round || round.status === "revealed";
  const showRevealVoteButton = round?.status === "voting";
  const showEndSessionButton = !round || round.status === "revealed";

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

  const handleEndSession = async () => {
    try {
      if (!session) {
        return;
      }
      const client = new ApiClient();
      await client.endSession(session.sessionId);
      onClick("endSession");
    } catch (error) {
      onError("endSession", error);
    }
  };

  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      {hasEnableSession && (
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
          {showEndSessionButton && (
            <button
              type="button"
              className="btn btn-error w-full"
              aria-label="セッションを終了"
              onClick={handleEndSession}
            >
              <StopCircle />
              セッションを終了
            </button>
          )}
        </div>
      )}
      {!hasEnableSession && (
        <div className="card-body bg-neutral-content text-left">
          <p>セッション終了済み</p>
        </div>
      )}
    </div>
  );
}
