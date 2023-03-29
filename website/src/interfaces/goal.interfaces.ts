export interface Goal {
    id?: number,
    goal_amount: number,
    name: string,
    slug: string,
    image: string,
    target_date: Date,
    accumulated_amount?: number,
    ads_by_amount?: number,
    catalog_url?: string,
    days_until_bang?: number,
}

export interface CreateUpdateGoal {
    id?: number,
    goal_amount: number,
    name: string,
    target_date: Date | undefined,
    image: string,
}