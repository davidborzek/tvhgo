import styles from './Modal.module.scss';
import { PropsWithChildren, useEffect, useRef } from 'react';
import ModalCloseButton from './ModalCloseButton';
import { c } from '@/utils/classNames';

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

  const handleClose = (event: React.MouseEvent<HTMLInputElement>) => {
    if (event.target === ref.current && !props.disableBackdropClose) {
      props.onClose();
    }
  };

  return (
    <div
      ref={ref}
      className={c(styles.wrapper, props.visible ? styles.visible : undefined)}
      onClick={handleClose}
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
