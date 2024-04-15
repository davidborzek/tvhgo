import { LoaderFunctionArgs, useLoaderData } from 'react-router-dom';

import { DVRConfig } from '@/clients/api/api.types';
import Headline from '@/components/common/headline/Headline';
import Pair from '@/components/common/pairList/Pair/Pair';
import PairKey from '@/components/common/pairList/PairKey/PairKey';
import PairList from '@/components/common/pairList/PairList';
import PairValue from '@/components/common/pairList/PairValue/PairValue';
import { getDVRConfig } from '@/clients/api/api';
import styles from './DVRConfigDetailView.module.scss';
import { useTranslation } from 'react-i18next';

// TODO: move to backend
const RETENTION_FOREVER = 2147483647;
const RETENTION_ON_FILE_REMOVAL = 2147483646;
const RETENTION_MAINTAINED_SPACE = 2147483646;

export function loader({ params }: LoaderFunctionArgs) {
  if (!params.id) {
    throw new Error('missing id parameter');
  }

  return getDVRConfig(params.id);
}

export function Component() {
  const { t } = useTranslation();
  const config = useLoaderData() as DVRConfig;

  const renderRetentionDays = () => {
    if (config.retentionDays === RETENTION_FOREVER) {
      return t('forever');
    }

    if (config.retentionDays === RETENTION_ON_FILE_REMOVAL) {
      return t('on_file_removal');
    }

    return t('days', { count: config.retentionDays });
  };

  const renderRemovalDays = () => {
    if (config.removalDays === RETENTION_FOREVER) {
      return t('forever');
    }

    if (config.removalDays === RETENTION_MAINTAINED_SPACE) {
      return t('maintained_space');
    }

    return t('days', { count: config.removalDays });
  };

  const renderPriority = () => {
    switch (config.priority) {
      case 'important':
        return t('important');
      case 'high':
        return t('high');
      case 'normal':
        return t('normal');
      case 'low':
        return t('low');
      case 'unimportant':
        return t('unimportant');
      case 'default':
        return t('default');
    }

    return t('unknown');
  };

  return (
    <div className={styles.view}>
      <h1>{config.name || t('default_profile')}</h1>
      <div className={styles.row}>
        <div className={styles.section}>
          <Headline>{t('common')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('enabled')}</PairKey>
              <PairValue>{config.enabled ? t('yes') : t('no')}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('priority')}</PairKey>
              <PairValue>{renderPriority()}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('retention')}</PairKey>
              <PairValue>{renderRetentionDays()}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('removal')}</PairKey>
              <PairValue>{renderRemovalDays()}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('clone_schedule_on_error')}</PairKey>
              <PairValue>{config.clone ? t('yes') : t('no')}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('rerecord_errors')}</PairKey>
              <PairValue>{config.rerecordErrors}</PairValue>
            </Pair>
          </PairList>
        </div>

        <div className={styles.section}>
          <Headline>{t('recording_time')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('recording_minutes_before_start')}</PairKey>
              <PairValue>
                {t('minutes', { value: config.startPadding })}
              </PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('recording_minutes_after_end')}</PairKey>
              <PairValue>
                {t('minutes', { value: config.endPadding })}
              </PairValue>
            </Pair>
          </PairList>
        </div>
      </div>

      <div className={styles.row}>
        <div className={styles.section}>
          <Headline>{t('storage')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('path')}</PairKey>
              <PairValue>{config.storage.path}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('maintain_free_space')}</PairKey>
              <PairValue>{config.storage.maintainFreeSpace} MB</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('maintain_used_space')}</PairKey>
              <PairValue>{config.storage.maintainUsedSpace} MB</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('directory_permissions')}</PairKey>
              <PairValue>{config.storage.directoryPermissions}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('file_permissions')}</PairKey>
              <PairValue>{config.storage.filePermissions}</PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('charset')}</PairKey>
              <PairValue>{config.storage.charset}</PairValue>
            </Pair>
          </PairList>
        </div>

        <div className={styles.section}>
          <Headline>{t('subdirectories')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('channel_subdir')}</PairKey>
              <PairValue>
                {config.subdirectories.channelSubdir ? t('yes') : t('no')}
              </PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('day_subdir')}</PairKey>
              <PairValue>
                {config.subdirectories.daySubdir ? t('yes') : t('no')}
              </PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('title_subdir')}</PairKey>
              <PairValue>
                {config.subdirectories.titleSubdir ? t('yes') : t('no')}
              </PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('tvmovies_subdir_format')}</PairKey>
              <PairValue>
                {config.subdirectories.tvMoviesSubdirFormat}
              </PairValue>
            </Pair>

            <Pair>
              <PairKey>{t('tvshows_subdir_format')}</PairKey>
              <PairValue>{config.subdirectories.tvShowsSubdirFormat}</PairValue>
            </Pair>
          </PairList>
        </div>
      </div>

      <div className={styles.section}>
        <Headline>{t('file')}</Headline>
        <div className={styles.row}>
          <div className={styles.section}>
            <PairList>
              <Pair>
                <PairKey>{t('include_channel')}</PairKey>
                <PairValue>
                  {config.file.includeChannel ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('include_date')}</PairKey>
                <PairValue>
                  {config.file.includeDate ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('include_time')}</PairKey>
                <PairValue>
                  {config.file.includeTime ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('include_episode')}</PairKey>
                <PairValue>
                  {config.file.includeEpisode ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('include_subtitle')}</PairKey>
                <PairValue>
                  {config.file.includeSubtitle ? t('yes') : t('no')}
                </PairValue>
              </Pair>
            </PairList>
          </div>
          <div className={styles.section}>
            <PairList>
              <Pair>
                <PairKey>{t('omit_title')}</PairKey>
                <PairValue>
                  {config.file.omitTitle ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('clean_title')}</PairKey>
                <PairValue>
                  {config.file.cleanTitle ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('allow_whitespace')}</PairKey>
                <PairValue>
                  {config.file.allowWhitespace ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('windows_compatible_filename')}</PairKey>
                <PairValue>
                  {config.file.windowsCompatibleFilename ? t('yes') : t('no')}
                </PairValue>
              </Pair>

              <Pair>
                <PairKey>{t('tag_files')}</PairKey>
                <PairValue>
                  {config.file.tagFiles ? t('yes') : t('no')}
                </PairValue>
              </Pair>
            </PairList>
          </div>
        </div>
      </div>

      <div className={styles.row}>
        <div className={styles.section}>
          <Headline>{t('epg')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('duplicate_handling')}</PairKey>
              <PairValue>{config.epg.duplicateHandling}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('epg_update_window')}</PairKey>
              <PairValue>{config.epg.epgUpdateWindow}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('epg_running')}</PairKey>
              <PairValue>
                {config.epg.epgRunning ? t('yes') : t('no')}
              </PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('skip_commercials')}</PairKey>
              <PairValue>
                {config.epg.skipCommercials ? t('yes') : t('no')}
              </PairValue>
            </Pair>
          </PairList>
        </div>

        <div className={styles.section}>
          <Headline>{t('autorec')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('max_count')}</PairKey>
              <PairValue>{config.epg.autorec.maxCount}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('max_schedules')}</PairKey>
              <PairValue>{config.epg.autorec.maxSchedules}</PairValue>
            </Pair>
          </PairList>
        </div>
      </div>

      <div className={styles.row}>
        <div className={styles.section}>
          <Headline>{t('artwork')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('fetch')}</PairKey>
              <PairValue>{config.artwork.fetch ? t('yes') : t('no')}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('allow_unidentifiable_broadcasts')}</PairKey>
              <PairValue>
                {config.artwork.allowUnidentifiableBroadcasts
                  ? t('yes')
                  : t('no')}
              </PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('command_line_options')}</PairKey>
              <PairValue>{config.artwork.commandLineOptions || '-'}</PairValue>
            </Pair>
          </PairList>
        </div>

        <div className={styles.section}>
          <Headline>{t('hooks')}</Headline>
          <PairList>
            <Pair>
              <PairKey>{t('start')}</PairKey>
              <PairValue>{config.hooks.start || '-'}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('stop')}</PairKey>
              <PairValue>{config.hooks.stop || '-'}</PairValue>
            </Pair>
            <Pair>
              <PairKey>{t('remove')}</PairKey>
              <PairValue>{config.hooks.remove || '-'}</PairValue>
            </Pair>
          </PairList>
        </div>
      </div>
    </div>
  );
}

Component.displayName = 'DVRConfigDetailView';
