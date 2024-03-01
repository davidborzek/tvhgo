import ChannelSelectModal from '@/components/channels/selectModal/ChannelSelectModal';

import styles from './RecordingCreateView.module.scss';
import { useState } from 'react';
import { Channel } from '@/clients/api/api.types';
import Button from '@/components/common/button/Button';

const RecordingCreateView = () => {
  const [channel, setChannel] = useState<Channel | null>(null);
  const [channelSelectionVisible, setChannelSelectionVisible] = useState(false);

  return (
    <>
      <ChannelSelectModal
        channels={[...new Array(50)].map((_, index) => ({
          enabled: true,
          id: `ID ${index}`,
          name: `Channel ${index}`,
          number: index,
          piconId: index,
        }))}
        onClose={() => {
          setChannelSelectionVisible(false);
        }}
        visible={channelSelectionVisible}
        onSelect={(ch) => {
          setChannel(ch);
          setChannelSelectionVisible(false);
        }}
      />

      <Button label="Open" onClick={() => setChannelSelectionVisible(true)} />
      {channel?.name}
    </>
  );
};

export default RecordingCreateView;
