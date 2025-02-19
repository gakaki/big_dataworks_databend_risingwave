import React, { useState } from 'react';
import './StarRating.css';

const StarRating: React.FC = () => {
  const [rating, setRating] = useState(0);
  const [hover, setHover] = useState(0);

  return (
    <div className="star-rating">
      {[1, 2, 3, 4, 5].map((star) => (
        <button
          type="button"
          key={star}
          className={`star-button ${ (hover || rating) >= star ? 'on' : 'off' }`}
          onClick={() => setRating(star)}
          onMouseEnter={() => setHover(star)}
          onMouseLeave={() => setHover(0)}
        >
          <svg
            className="star-svg"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path d="M10 15l-5.878 3.09 1.122-6.545L.488 7.91l6.561-.955L10 1l2.951 5.955 6.561.955-4.756 4.635 1.122 6.545z"/>
          </svg>
        </button>
      ))}
    </div>
  );
};

export default StarRating;