import { useCallback, useEffect, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import type { RoundParticipant } from "../types/Participant";
import type { Round } from "../types/Round";
import { type Session, isSessionStatus } from "../types/Session";
import type { Vote } from "../types/Vote";

export type ReturnValue = {
  session: Session | null;
  round: Round | null;
  myVote: Vote | null;
  error: unknown | null;
  reload: () => Promise<void>;
};

export default function useSession(
  sessionId: string,
  participantId: string,
): ReturnValue {
  const [session, setSession] = useState<Session | null>(null);
  const [round, setRound] = useState<Round | null>(null);
  const [myVote, setMyVote] = useState<Vote | null>(null);
  const [error, setError] = useState<unknown | null>(null);

  useEffect(() => {
    (async () => {
      await reload();
    })();
  }, []);

  const reload = useCallback(async () => {
    try {
      const apiClient = new ApiClient();

      const session = await fetchSession(apiClient, sessionId);
      setSession(session);

      if (!session) {
        setRound(null);
        setMyVote(null);
        setError(null);
        return;
      }

      const round = await fetchRound(apiClient, session, participantId);
      setRound(round);

      if (!round) {
        setMyVote(null);
        setError(null);
        return;
      }

      setMyVote(findMyVote(participantId, round));
      setError(null);
    } catch (error) {
      setError(error);
    }
  }, [sessionId, participantId]);

  const fetchSession = async (
    apiClient: ApiClient,
    sessionId: string,
  ): Promise<Session | null> => {
    const response = await apiClient.fetchSession(sessionId);
    if (!response.session) {
      return null;
    }
    if (!isSessionStatus(response.session.status)) {
      throw new Error(
        `Invalid session status, status=${response.session.status}`,
      );
    }

    const participants = response.session.participants.map((participant) => {
      return {
        id: participant.participantId,
        name: participant.name,
      };
    });

    return {
      sessionId: response.session.sessionId,
      sessionName: response.session.sessionName,
      participants,
      currentRoundId: response.session.currentRoundId,
      hostId: response.session.hostId,
      status: response.session.status,
    };
  };

  const fetchRound = async (
    apiClient: ApiClient,
    session: Session,
    participantId: string,
  ): Promise<Round | null> => {
    if (!session.currentRoundId) {
      return null;
    }
    const response = await apiClient.fetchRound(
      session.currentRoundId,
      participantId,
    );
    if (!response.round) {
      return null;
    }
    const round = response.round;

    const votes = round.votes.map((vote) => {
      return {
        participantId: vote.participantId,
        vote: vote.value ?? null,
      };
    });
    const participants: RoundParticipant[] = session.participants.map(
      (participant) => {
        const isVoted = round.votes.some(
          (vote) => vote.participantId === participant.id,
        );
        const vote = round.votes.find(
          (vote) => vote.participantId === participant.id,
        );

        return {
          id: participant.id,
          name: participant.name,
          isVoted,
          vote: vote?.value ?? null,
        };
      },
    );

    return {
      ...response.round,
      votes,
      participants,
    };
  };

  const findMyVote = (participantId: string, round: Round): Vote | null => {
    const myVote = round.votes.find(
      (vote) => vote.participantId === participantId,
    );
    return myVote ?? null;
  };

  return {
    session,
    round,
    myVote,
    error,
    reload,
  };
}
