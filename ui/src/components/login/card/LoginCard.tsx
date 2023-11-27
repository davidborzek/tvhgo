import React, { PropsWithChildren } from 'react';

import styles from './LoginCard.module.scss';

type Props = {
  title?: string;
  onSubmit?: React.FormEventHandler<HTMLFormElement>;
};

function LoginCard(props: PropsWithChildren<Props>) {
  return (
    <div className={styles.card}>
      <div className={styles.imageHeader}>
        <img className={styles.image} src="/img/tvhgo.png" />
      </div>
      <form
        title={props.title}
        className={styles.form}
        onSubmit={props.onSubmit}
      >
        {props.children}
      </form>
    </div>
  );
}

export default LoginCard;
