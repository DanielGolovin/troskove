export type Category = {
    id: string;
    name: string;
}

export type CategoryWithTotal = Category & {
    total: number;
}
