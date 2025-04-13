export interface SessionParticipant {
  id: string;
  name: string;
}

export interface RoundParticipant extends SessionParticipant {
  isVoted: boolean;
  vote: string | number | null;
}
