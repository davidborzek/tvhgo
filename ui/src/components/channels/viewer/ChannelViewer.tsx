import { EpgEvent, StreamProfile } from '@/clients/api/api.types';
import React, { useEffect, useState } from 'react';
import { getChannelStreamUrl, getStreamProfiles } from '@/clients/api/api';

import Button from '@/components/common/button/Button';
import Dropdown from '@/components/common/dropdown/Dropdown';
import VideoPlayer from '@/components/common/videoPlayer/VideoPlayer';
import styles from './ChannelViewer.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  event: EpgEvent;
  isOpen: boolean;
  onClose: () => void;
};

function ChannelViewer({ event, isOpen, onClose }: Props) {
  const { t } = useTranslation();
  const [profiles, setProfiles] = useState<StreamProfile[]>([]);
  const [selectedProfile, setSelectedProfile] = useState<string>('');
  const [streamUrl, setStreamUrl] = useState<string>('');

  useEffect(() => {
    const loadProfiles = async () => {
      try {
        const streamProfiles = await getStreamProfiles();
        setProfiles(streamProfiles);
        // Set default profile if available
        if (streamProfiles.length > 0) {
          setSelectedProfile(streamProfiles[0].id);
        }
      } catch (error) {
        console.error('Failed to load stream profiles:', error);
      }
    };

    if (isOpen) {
      loadProfiles();
    }
  }, [isOpen]);

  useEffect(() => {
    if (isOpen && selectedProfile) {
      const url = getChannelStreamUrl(event.channelNumber, selectedProfile);
      setStreamUrl(url);
    }
  }, [isOpen, selectedProfile, event.channelNumber]);

  const handleProfileChange = (profileId: string) => {
    setSelectedProfile(profileId);
  };

  const handleClose = () => {
    setStreamUrl('');
    onClose();
  };

  if (!isOpen) {
    return null;
  }

  return (
    <div className={styles.overlay}>
      <div className={styles.container}>
        <div className={styles.header}>
          <div className={styles.channelInfo}>
            <h3>{event.channelName}</h3>
            <p>{event.title}</p>
          </div>
          <div className={styles.controls}>
            <Dropdown
              value={selectedProfile}
              onChange={handleProfileChange}
              options={profiles.map((profile) => ({
                value: profile.id,
                title: profile.name,
              }))}
              label={t('select_stream_profile')}
            />
            <Button onClick={handleClose} style="text" label={t('close')} />
          </div>
        </div>
        <div className={styles.videoContainer}>
          {streamUrl && (
            <VideoPlayer
              src={streamUrl}
              onError={(error) => console.error('Video error:', error)}
            />
          )}
        </div>
      </div>
    </div>
  );
}

export default ChannelViewer;
