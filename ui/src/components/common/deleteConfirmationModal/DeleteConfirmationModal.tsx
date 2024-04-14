import Button from '@/components/common/button/Button';
import Modal from '@/components/common/modal/Modal';
import { TestIds } from '@/__test__/ids';
import styles from './DeleteConfirmationModal.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  title?: string | null | undefined;
  buttonTitle?: string | null | undefined;
  visible: boolean;
  onClose: () => void;
  onConfirm: () => void;
  pending?: boolean;
};

const DeleteConfirmationModal = ({
  visible,
  onClose,
  onConfirm,
  title,
  buttonTitle,
  pending,
}: Props) => {
  const { t } = useTranslation();

  return (
    <Modal onClose={onClose} visible={visible}>
      <div className={styles.content}>
        {title ? <h3 className={styles.headline}>{title}</h3> : <></>}
        <Button
          disabled={pending}
          label={buttonTitle || t('delete')}
          style="red"
          onClick={onConfirm}
          testID={TestIds.CONFIRM_DELETE_BUTTON}
        />
      </div>
    </Modal>
  );
};

export default DeleteConfirmationModal;
