import { FC } from "react";
import styles from "./index.module.css";

export type TransactionRowItemProps = {
  id: string;
  amount: number;
  date: string;
  category: string;
};

export const TransactionRowItem: FC<TransactionRowItemProps> = ({
  amount,
  date,
  category,
}) => {
  return (
    <div className={styles.container}>
      <div>{amount}</div>
      <div>
        <div style={{ textAlign: "right" }}>{category}</div>
        <div>{date}</div>
      </div>
    </div>
  );
};
