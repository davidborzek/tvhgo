import { NavLink, useLoaderData } from 'react-router-dom';

import Badge from '@/components/common/badge/Badge';
import Button from '@/components/common/button/Button';
import { DVRConfig } from '@/clients/api/api.types';
import Headline from '@/components/common/headline/Headline';
import Table from '@/components/common/table/Table';
import TableBody from '@/components/common/table/TableBody';
import TableCell from '@/components/common/table/TableCell';
import TableHead from '@/components/common/table/TableHead';
import TableHeadCell from '@/components/common/table/TableHeadCell';
import TableRow from '@/components/common/table/TableRow';
import Tooltip from '@/components/common/tooltip/Tooltip';
import { getDVRConfigs } from '@/clients/api/api';
import styles from './DVRConfigListView.module.scss';
import { useDVRConfig } from '@/hooks/dvr';
import { useTranslation } from 'react-i18next';

export function loader() {
  return getDVRConfigs();
}

export function Component() {
  const { t } = useTranslation();
  const configs = useLoaderData() as DVRConfig[];
  const [isPending, deleteDVRConfig] = useDVRConfig();

  return (
    <div className={styles.view}>
      <Headline>{t('dvr_profiles')}</Headline>
      <div>
        <Table>
          <TableHead>
            <TableHeadCell>{t('name')}</TableHeadCell>
            <TableHeadCell className={styles.path}>{t('path')}</TableHeadCell>
            <TableHeadCell className={styles.time}>
              {t('recording_minutes_before_start')}
            </TableHeadCell>
            <TableHeadCell className={styles.time}>
              {t('recording_minutes_after_end')}
            </TableHeadCell>
            <TableHeadCell className={styles.default}>
              {t('default')}
            </TableHeadCell>
            <TableHeadCell />
          </TableHead>
          <TableBody>
            {configs.map((config) => (
              <TableRow>
                <TableCell>
                  <NavLink
                    to={`/dvr/config/${config.id}`}
                    className={styles.link}
                  >
                    {config.name || t('default_profile')}
                  </NavLink>
                </TableCell>
                <TableCell className={styles.path}>
                  {config.storage.path}
                </TableCell>
                <TableCell className={styles.time}>
                  {t('minutes', { value: config.startPadding })}
                </TableCell>
                <TableCell className={styles.time}>
                  {t('minutes', { value: config.endPadding })}
                </TableCell>
                <TableCell className={styles.default}>
                  {config.original ? (
                    <Badge color="default">{t('yes')}</Badge>
                  ) : (
                    <Badge color="failure">{t('no')}</Badge>
                  )}
                </TableCell>
                <TableCell className={styles.actions}>
                  <Tooltip
                    className={styles.deleteTooltip}
                    direction="left"
                    text={t('dvr_delete_default_profile_info')}
                    disabled={!config.original}
                  >
                    <Button
                      disabled={config.original || isPending}
                      loading={isPending}
                      style="red"
                      size="small"
                      label={t('delete')}
                      quiet
                      onClick={() => deleteDVRConfig(config.id)}
                    />
                  </Tooltip>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}

Component.displayName = 'DVRConfigListView';
