export type TransactionType = "expenses" | "mixed" | "incomes";

export type Transaction = {
  id: string;
  amount: number;
  date: string;
  category: string;
};

export type TransactionMonth = {
  total: number;
  name: string;
};

export type TransactionYear = {
  name: string;
  total: number;
};
