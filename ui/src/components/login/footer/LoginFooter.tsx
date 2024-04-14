import styles from './LoginFooter.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  commitHash: string;
  githubUrl: string;
  version: string;
};

function LoginFooter(props: Props) {
  const { t } = useTranslation();

  return (
    <footer className={styles.footer}>
      <span>{t('made_with')}</span>
      <span className={styles.sep}> | </span>
      <span>
        <a
          className={styles.link}
          href={props.githubUrl}
          target="_blank"
          rel="noopener noreferrer"
        >
          Open Source
        </a>
      </span>
      <span className={styles.sep}> | </span>
      <span>
        <a
          className={styles.link}
          href={`${props.githubUrl}/commit/${props.commitHash}`}
          target="_blank"
          rel="noopener noreferrer"
        >
          {`${props.version} (${props.commitHash})`}
        </a>
      </span>
    </footer>
  );
}

export default LoginFooter;
