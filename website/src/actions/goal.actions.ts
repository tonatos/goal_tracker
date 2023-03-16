import { useSetRecoilState } from "recoil";
import { useFetchWrapper } from "../helpers";
import { currentGoalAtom, goalsAtom } from "../recoil/goal";

export const useGoalActions = () => {
    const baseUrl = `${process.env.REACT_APP_API_URL}/goal`;
    const fetchWrapper = useFetchWrapper();
    const setGoals = useSetRecoilState(goalsAtom);
    const setCurrentGoal = useSetRecoilState(currentGoalAtom);

    const getAll = async () => {
        const r = await fetchWrapper.get(baseUrl, null);
        return setGoals(r.data)
    }

    const getDefault =  async (id: number | undefined) => {
        if (!id) {
            return;
        }
        const r = await fetchWrapper.get(`${baseUrl}/${id}`, null);
        return setCurrentGoal(r.data);
    }
    
    return {
        getAll,
        getDefault,
    }
}