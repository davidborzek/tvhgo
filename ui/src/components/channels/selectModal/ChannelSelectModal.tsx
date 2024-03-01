import Input from '@/components/common/input/Input';
import Modal from '@/components/common/modal/Modal';
import styles from './ChannelSelectModal.module.scss';
import { useTranslation } from 'react-i18next';
import { Channel } from '@/clients/api/api.types';
import { useRef, useState } from 'react';
import { useGetChannels } from '@/hooks/channel';
import { useNavigate } from 'react-router-dom';
import { useDebounce } from '@/hooks/debounce';

const ChannelSelectModal = () => {
  const ref = useRef<HTMLDivElement>(null);
  const { t } = useTranslation();
  const navigate = useNavigate();

  const [searchValue, setSearchValue] = useState('');
  const { channels, error } = useGetChannels(useDebounce(searchValue));

  const handleClose = (ch?: Channel) => {
    navigate('/recordings/create', { state: { channel: ch } });

    setSearchValue('');
    if (ref.current) {
      ref.current.scrollTop = 0;
    }
  };

  const renderChannel = (channel: Channel) => {
    return (
      <div className={styles.channel} onClick={() => handleClose(channel)}>
        <div className={styles.piconContainer}>
          <img className={styles.picon} src={`/api/picon/${channel.piconId}`} />
        </div>
        <span>{channel.name}</span>
      </div>
    );
  };

  return (
    <Modal
      onClose={() => {
        handleClose();
      }}
      visible
    >
      <h3 className={styles.headline}>{t('select_channel')}</h3>
      <Input
        value={searchValue}
        onChange={(e) => setSearchValue(e.target.value)}
        placeholder={t('search')}
        fullWidth
      />
      <div ref={ref} className={styles.channels}>
        {channels.map(renderChannel)}
      </div>
    </Modal>
  );
};

export default ChannelSelectModal;
