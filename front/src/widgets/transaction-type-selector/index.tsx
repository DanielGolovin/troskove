import { useState } from "react";
import { TransactionTypeIndicators } from "./ui/transaction-type-indicator";
import { TransactionsTypeSelector } from "./ui/transaction-type-selector";

import styles from "./index.module.css";
import { Category } from "../../entities/categories";
import { TransactionType } from "../../entities/transactions";
import { useGetMonthTotalQuery } from "./data-hooks/use-get-month-total-query";

const TransactionTypeToIndicatorMap: Record<TransactionType, number> = {
  expenses: 1,
  mixed: 2,
  incomes: 3,
};

const transactionTypes = Object.keys(
  TransactionTypeToIndicatorMap
) as TransactionType[];

const amount = Object.keys(TransactionTypeToIndicatorMap).length;

type TransactionTypeSelectorProps = {
  excludedCategoryIds: Category["id"][];
  onTransactionTypeChange: (type: TransactionType) => void;
};

export const TransactionTypeSelector: React.FC<
  TransactionTypeSelectorProps
> = ({ onTransactionTypeChange, excludedCategoryIds }) => {
  const [selectedTransactionType, setSelectedTransactionType] =
    useState<TransactionType>("mixed");

  const { data: total = 0 } = useGetMonthTotalQuery({
    transactionType: selectedTransactionType,
    excludedCategoryIds,
  });

  const onIndicatorClick = (index: number) => {
    const selectedTransactionType = transactionTypes[index];
    setSelectedTransactionType(selectedTransactionType);
    onTransactionTypeChange(selectedTransactionType);
  };

  const selectedTransactionTypeIndicator =
    TransactionTypeToIndicatorMap[selectedTransactionType];

  return (
    <header className={styles.container}>
      <TransactionTypeIndicators
        amount={amount}
        selected={selectedTransactionTypeIndicator}
        onClick={onIndicatorClick}
      />
      <TransactionsTypeSelector
        total={total + "$"}
        selected={selectedTransactionType}
        types={transactionTypes}
        onSelect={setSelectedTransactionType}
      />
    </header>
  );
};
