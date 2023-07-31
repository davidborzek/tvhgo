import styles from './Modal.module.scss';
import { PropsWithChildren, useEffect, useRef } from 'react';
import ModalCloseButton from './ModalCloseButton';

type Props = {
  visible: boolean;
  onClose: () => void;
  disableBackdropClose?: boolean;
  maxWidth?: string | number;
};

export default function Modal(props: PropsWithChildren<Props>) {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      e.key == 'Escape' && props.onClose();
    };

    document.addEventListener('keydown', handleKeyDown);

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, []);

  return (
    <div
      ref={ref}
      className={`${styles.modalWrapper} ${
        props.visible ? styles.visible : ''
      }`}
      onClick={(event) => {
        if (event.target === ref.current) {
          !props.disableBackdropClose && props.onClose();
        }
      }}
    >
      <div className={styles.modal} style={{ maxWidth: props.maxWidth }}>
        <div className={styles.container}>
          <ModalCloseButton onClick={props.onClose} />
          {props.children}
        </div>
      </div>
    </div>
  );
}
