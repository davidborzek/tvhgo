import { c } from '@/utils/classNames';
import { ButtonStyle, getStyleClass } from './Button';
import styles from './Button.module.scss';

type Props = {
  label: string;
  download?: string | boolean;
  href?: string;
  className?: string;
  style?: ButtonStyle;
};

const ButtonLink = (props: Props) => {
  return (
    <a
      className={c(
        styles.button,
        getStyleClass(props.style),
        props.className ? props.className : ''
      )}
      href={props.href}
      download={props.download}
    >
      {props.label}
    </a>
  );
};

export default ButtonLink;
