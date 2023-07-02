import styles from './Modal.module.scss';
import { PropsWithChildren, useRef } from 'react';
import ModalCloseButton from './ModalCloseButton';

type Props = {
  visible: boolean;
  onClose: () => void;
};

export default function Modal(props: PropsWithChildren<Props>) {
  const ref = useRef<HTMLDivElement>(null);

  return (
    <div
      ref={ref}
      className={`${styles.modalWrapper} ${
        props.visible ? styles.visible : ''
      }`}
      onClick={(event) => {
        if (event.target === ref.current) {
          props.onClose();
        }
      }}
    >
      <div className={styles.modal}>
        <div className={styles.container}>
          <ModalCloseButton onClick={props.onClose} />
          {props.children}
        </div>
      </div>
    </div>
  );
}
