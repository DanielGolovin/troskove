import { TransactionYearRow } from "./ui/transaction-year-row";
import { TransactionMonthRow } from "./ui/transaction-month-row";
import { TransactionRowItem } from "./ui/transaction-row";

export const TransactionsList = () => {
  return (
    <div>
      <TransactionYearRow
        amount={100}
        date={"2024"}
        MonthRow={TransactionMonthRow}
        TransactionRow={TransactionRowItem}
      />
      <TransactionYearRow
        amount={100}
        date={"2023"}
        MonthRow={TransactionMonthRow}
        TransactionRow={TransactionRowItem}
      />
    </div>
  );
};
