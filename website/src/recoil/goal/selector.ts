import { selector } from "recoil";
import { currentGoalAtom, goalsAtom } from "./atom"

export const goalsListQuery = selector({
    key: 'goalsListQuery',
    get: ({ get }) => get(goalsAtom) 
});

export const goalInfoQuery = selector({
    key: 'goalInfoQueryCheck', 
    get: ({ get }) => get(currentGoalAtom)
});
