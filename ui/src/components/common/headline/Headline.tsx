import { PropsWithChildren } from 'react';
import styles from './Headline.module.scss';

const Headline = (props: PropsWithChildren) => {
  return <h2 className={styles.headline}>{props.children}</h2>;
};

export default Headline;
