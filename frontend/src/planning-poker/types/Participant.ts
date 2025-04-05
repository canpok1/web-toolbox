export interface SessionParticipant {
  id: string;
  name: string;
}

export interface RoundParticipant extends SessionParticipant {
  vote: string | number | null;
}
