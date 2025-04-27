import { useCallback, useEffect, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import type { SessionParticipant } from "../types/Participant";
import type { Round } from "../types/Round";
import { type Session, isSessionStatus } from "../types/Session";
import type { Vote } from "../types/Vote";

export type ReturnValue = {
  session: Session | null;
  round: Round | null;
  myVote: Vote | null;
  myParticipant: SessionParticipant | null;
  loaded: boolean;
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
  const [myParticipant, setMyParticipant] = useState<SessionParticipant | null>(
    null,
  );
  const [loaded, setLoaded] = useState<boolean>(false);
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

      setMyParticipant(
        session.participants.find(
          (participant) => participant.id === participantId,
        ) ?? null,
      );

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
    } finally {
      setLoaded(true);
    }
  }, [sessionId, participantId]);

  const fetchSession = async (
    apiClient: ApiClient,
    sessionId: string,
  ): Promise<Session | null> => {
    const response = await apiClient.fetchSession(sessionId);
    if (!response) {
      throw new Error("failed to fetchSession, response is null");
    }
    if (!response.session) {
      throw new Error(
        "failed to fetchSession, response has not session property",
      );
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
      participants,
      currentRoundId: response.session.currentRoundId,
      hostId: response.session.hostId,
      status: response.session.status,
      scaleType: response.session.scaleType,
      scales: response.session.scales,
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
        participantName: vote.participantName,
        vote: vote.value ?? null,
      };
    });

    return {
      ...response.round,
      votes,
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
    myParticipant,
    loaded,
    error,
    reload,
  };
}
