import { useQuery } from '@tanstack/react-query'
import { TransactionType } from 'src/entities/transactions';

export const getCategoriesQueryKeyBase = 'month-total'

type Options = {
  transactionType: TransactionType;
  excludedCategoryIds?: string[];
}

export const useGetMonthTotalQuery = ({
  transactionType,
  excludedCategoryIds = []
}: Options) => {
  const { data } = useQuery<number>({
    queryKey: [getCategoriesQueryKeyBase, transactionType, ...excludedCategoryIds],
    queryFn: () => Promise.resolve(mockData[transactionType]),
    // queryFn: () => fetch('/api/categories').then((res) => res.json()),
  });

  return { data };
}

const mockData = {
  expenses: -1500,
  incomes: 1200,
  mixed: -300,
};
  