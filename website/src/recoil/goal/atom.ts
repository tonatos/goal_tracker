import { atom } from "recoil"
import { Goal } from "../../interfaces/goal.interfaces";

export const goalsAtom = atom<Goal[]>({
    key: "goalsAtom",
    default: [] as Goal[],
});
export const currentGoalAtom = atom<Goal>({
    key: "currentGoalAtom",
    default: {} as Goal,
});