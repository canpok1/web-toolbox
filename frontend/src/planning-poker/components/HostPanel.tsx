import { CheckCircle2, Play, StopCircle } from "lucide-react";
import { ApiClient } from "../../api/ApiClient";
import type { RoundStatus } from "../types/Round";

export type HostPanelEvent = "startRound" | "revealVotes" | "endSession";

export type HostPanelProps = {
  sessionId: string;
  roundId: string | undefined;
  roundStatus: RoundStatus | undefined;
  onClick: (event: HostPanelEvent) => void;
  onError: (event: HostPanelEvent, error: unknown) => void;
};

export default function HostPanel({
  sessionId,
  roundId,
  roundStatus,
  onClick,
  onError,
}: HostPanelProps) {
  const handleStartRound = async () => {
    try {
      const client = new ApiClient();
      await client.startRound(sessionId);
      onClick("startRound");
    } catch (error) {
      onError("startRound", error);
    }
  };

  const handleRevealVotes = async () => {
    try {
      if (roundId) {
        const client = new ApiClient();
        await client.revealRound(roundId);
        onClick("revealVotes");
      }
    } catch (error) {
      onError("revealVotes", error);
    }
  };

  const handleEndSession = async () => {
    try {
      const client = new ApiClient();
      await client.endSession(sessionId);
      onClick("endSession");
    } catch (error) {
      onError("endSession", error);
    }
  };

  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        {roundStatus !== "voting" && (
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
        {roundStatus === "voting" && roundId && (
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
        {roundStatus !== "voting" && (
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
    </div>
  );
}
