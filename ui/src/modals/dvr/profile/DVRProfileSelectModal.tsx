import Modal, { ModalProps } from '@/components/common/modal/Modal';

import Button from '@/components/common/button/Button';
import { DVRConfig } from '@/clients/api/api.types';
import Dropdown from '@/components/common/dropdown/Dropdown';
import styles from './DVRProfileSelectModal.module.scss';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

type Props = ModalProps & {
  profiles: Array<DVRConfig>;
  handleCreateRecording: (profileId: string) => void;
};

const DVRProfileSelectModal = ({
  profiles,
  handleCreateRecording,
  ...rest
}: Props) => {
  const { t } = useTranslation();

  const [profile, setProfile] = useState(profiles[0].id);

  return (
    <Modal {...rest}>
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('create_recording')}</h3>

        <div className={styles.form}>
          <Dropdown
            options={profiles.map((profile) => ({
              title: profile.name || t('default_profile'),
              value: profile.id,
            }))}
            value={profile}
            onChange={setProfile}
            label={t('profile')}
            fullWidth
          />

          <Button
            label={t('create')}
            onClick={() => {
              handleCreateRecording(profile);
              rest.onClose();
            }}
          />
        </div>
      </div>
    </Modal>
  );
};

export default DVRProfileSelectModal;
