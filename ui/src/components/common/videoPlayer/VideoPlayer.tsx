import React, { useEffect, useRef, useState } from 'react';

import styles from './VideoPlayer.module.scss';

type Props = {
  src: string;
  className?: string;
  onError?: (error: string) => void;
  onLoadStart?: () => void;
  onCanPlay?: () => void;
};

function VideoPlayer({
  src,
  className,
  onError,
  onLoadStart,
  onCanPlay,
}: Props) {
  const videoRef = useRef<HTMLVideoElement>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const video = videoRef.current;
    if (!video) return;

    const handleLoadStart = () => {
      setIsLoading(true);
      setError(null);
      onLoadStart?.();
    };

    const handleCanPlay = () => {
      setIsLoading(false);
      onCanPlay?.();
    };

    const handleError = () => {
      setIsLoading(false);
      const errorMessage = 'Failed to load video stream';
      setError(errorMessage);
      onError?.(errorMessage);
    };

    video.addEventListener('loadstart', handleLoadStart);
    video.addEventListener('canplay', handleCanPlay);
    video.addEventListener('error', handleError);

    return () => {
      video.removeEventListener('loadstart', handleLoadStart);
      video.removeEventListener('canplay', handleCanPlay);
      video.removeEventListener('error', handleError);
    };
  }, [onError, onLoadStart, onCanPlay]);

  useEffect(() => {
    const video = videoRef.current;
    if (video && src) {
      video.src = src;
      video.load();
    }
  }, [src]);

  return (
    <div className={`${styles.container} ${className || ''}`}>
      {isLoading && (
        <div className={styles.loading}>
          <div className={styles.spinner}></div>
          <span>Loading stream...</span>
        </div>
      )}
      {error && (
        <div className={styles.error}>
          <span>{error}</span>
        </div>
      )}
      <video
        ref={videoRef}
        className={styles.video}
        controls
        autoPlay
        muted
        playsInline
      />
    </div>
  );
}

export default VideoPlayer;
