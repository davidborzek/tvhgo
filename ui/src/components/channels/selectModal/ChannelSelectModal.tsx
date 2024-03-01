import Input from '@/components/common/input/Input';
import Modal from '@/components/common/modal/Modal';
import styles from './ChannelSelectModal.module.scss';
import { useTranslation } from 'react-i18next';
import { Channel } from '@/clients/api/api.types';
import { c } from '@/utils/classNames';

type Props = {
  channels: Channel[];
  visible: boolean;
  onClose: () => void;
  onSelect: (channel: Channel) => void;
};

const ChannelSelectModal = ({
  channels,
  onClose,
  onSelect,
  visible,
}: Props) => {
  const { t } = useTranslation();

  const renderChannel = (channel: Channel) => {
    return (
      <div className={styles.channel} onClick={() => onSelect(channel)}>
        <img src="picon" />
        <span>{channel.name}</span>
      </div>
    );
  };

  return (
    <Modal onClose={onClose} visible={visible}>
      <h3 className={styles.headline}>{t('select_channel')}</h3>
      <Input placeholder={t('search')} />
      <div className={styles.channels}>{channels.map(renderChannel)}</div>
    </Modal>
  );
};

export default ChannelSelectModal;
