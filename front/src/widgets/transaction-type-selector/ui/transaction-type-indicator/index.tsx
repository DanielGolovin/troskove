import clsx from "clsx";
import styles from "./index.module.css";

type TransactionTypeIndicatorsProps = {
  amount: number;
  selected: number;
  onClick?: (index: number) => void;
};

export const TransactionTypeIndicators: React.FC<
  TransactionTypeIndicatorsProps
> = ({ amount, selected, onClick }) => {
  const Items = Array.from({ length: amount }, (_, index) => {
    const handleClick = () => {
      onClick?.(index);
    };

    return (
      <div
        key={index}
        onClick={handleClick}
        className={clsx(styles.indicator, {
          [styles["indicator-selected"]]: index === selected - 1,
        })}
      ></div>
    );
  });

  return <nav className={styles["container"]}>{Items}</nav>;
};
