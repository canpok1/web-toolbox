import { useEffect, useRef, useState } from "react";
import { useParams, useSearchParams } from "react-router-dom";
import FibonacciVotePanel from "./components/FibonacciVotePanel";
import HostPanel, { type HostPanelEvent } from "./components/HostPanel";
import RoundSummary from "./components/RoundSummary";
import SessionSummary from "./components/SessionSummary";
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

  const [voteOption, setVoteOption] = useState<string | null>(null);
  const [roundParticipants, setRoundParticipants] = useState<
    RoundParticipant[]
  >([]);

  const intervalIdRef = useRef<NodeJS.Timeout | null>(null); // setIntervalのIDを保持

  // セッション情報を更新
  useEffect(() => {
    if (sessionId) {
      (async () => {
        await fetchSession(sessionId);
      })();
    }
  }, [sessionId, fetchSession]);

  // ラウンド情報を更新
  useEffect(() => {
    const roundId = session?.currentRoundId;
    if (roundId) {
      (async () => {
        await fetchRound(roundId, participandId);
      })();
    }
  }, [session, participandId, fetchRound]);

  // 自分の投票情報を更新
  useEffect(() => {
    const myVote = round?.votes.find(
      (vote) => vote.participantId === participandId,
    );
    setVoteOption(myVote?.vote ?? null);
  }, [round, participandId]);

  // ラウンド参加者情報を更新
  useEffect(() => {
    if (!session || !round) {
      return;
    }
    const participants: RoundParticipant[] = [];
    for (const participant of session.participants) {
      const isVoted = round.votes.some(
        (vote) => vote.participantId === participant.id,
      );
      const vote =
        round.votes.find((vote) => vote.participantId === participant.id)
          ?.vote ?? null;

      participants.push({
        id: participant.id,
        name: participant.name,
        isVoted,
        vote,
      });
    }

    setRoundParticipants(participants);
  }, [session, round]);

  // 定期更新
  useEffect(() => {
    intervalIdRef.current = setInterval(update, 5000);
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
  }, []);

  const update = async () => {
    if (sessionId) {
      await fetchSession(sessionId);
    }
    if (round?.roundId) {
      await fetchRound(round.roundId, participandId);
    }
  };

  const handleVote = (option: string) => {
    setVoteOption(option);
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
          <FibonacciVotePanel
            roundId={round.roundId}
            participantId={participandId}
            votedOption={voteOption}
            onAfterVote={handleVote}
          />
        )}

        {round && (
          <RoundSummary
            participants={roundParticipants}
            revealed={round?.status === "revealed"}
          />
        )}
      </div>
    </section>
  );
}

export default SessionPage;
