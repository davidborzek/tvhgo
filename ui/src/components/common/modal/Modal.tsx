import { PropsWithChildren, useEffect, useRef } from 'react';

import ModalCloseButton from './ModalCloseButton';
import styles from './Modal.module.scss';

type Props = {
  visible: boolean;
  onClose: () => void;
  disableBackdropClose?: boolean;
  disableEscapeClose?: boolean;
  maxWidth?: string | number;
};

export default function Modal({ onClose, ...props }: PropsWithChildren<Props>) {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape' && !props.disableEscapeClose) onClose();
    };

    document.addEventListener('keydown', handleKeyDown);

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [onClose]);

  return (
    <div
      ref={ref}
      className={`${styles.modalWrapper} ${
        props.visible ? styles.visible : ''
      }`}
      onClick={(event) => {
        if (event.target === ref.current) {
          if (!props.disableBackdropClose) onClose();
        }
      }}
    >
      <div
        className={styles.modal}
        style={{ maxWidth: props.maxWidth, width: '100%' }}
      >
        <div className={styles.container}>
          <ModalCloseButton onClick={onClose} />
          {props.children}
        </div>
      </div>
    </div>
  );
}
