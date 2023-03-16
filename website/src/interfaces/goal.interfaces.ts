export default interface Goal {
    id?: number,
    goal_amount: number,
    name: string,
    slug: string, 
    target_date: Date,
    accumulated_amount?: number,
    ads_by_amount?: number,
    catalog_url?: string,
    days_until_bang?: number,
}
