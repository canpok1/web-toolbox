import { ThumbsDown, ThumbsUp } from "lucide-react";
import type React from "react";

interface Props {
  liked: boolean;
  disliked: boolean;
  handleLikeClick: () => void;
  handleDislikeClick: () => void;
}

const FeedbackButtons: React.FC<Props> = ({
  liked,
  disliked,
  handleLikeClick,
  handleDislikeClick,
}) => {
  return (
    <div className="flex items-center justify-center gap-8">
      <button
        type="button"
        id="like-button"
        className={`feedback-button text-gray-600 hover:text-green-500 ${
          liked ? "liked" : ""
        }`}
        aria-label="良いテーマ"
        onClick={handleLikeClick}
      >
        <ThumbsUp className="h-6 w-6" />
      </button>
      <button
        type="button"
        id="dislike-button"
        className={`feedback-button text-gray-600 hover:text-red-500 ${
          disliked ? "disliked" : ""
        }`}
        aria-label="悪いテーマ"
        onClick={handleDislikeClick}
      >
        <ThumbsDown className="h-6 w-6" />
      </button>
    </div>
  );
};

export default FeedbackButtons;
