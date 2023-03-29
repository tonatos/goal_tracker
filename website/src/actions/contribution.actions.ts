import { useRecoilState } from "recoil";
import { useFetchWrapper } from "../helpers";
import { Contribution } from "../interfaces/contribution.interfaces";
import { contributionsAtom } from "../recoil/contribution";

export const useContributionActions = () => {
    const baseUrl = `${process.env.REACT_APP_API_URL}/goal`;
    const fetchWrapper = useFetchWrapper();
    const [contributions, setContributions] = useRecoilState(contributionsAtom);

    const getAll = async (goalId: number) => {
        const contributions = await fetchWrapper.get(`${baseUrl}/${goalId}`);
        return setContributions(contributions);
    }

    const create = async (goalId: number | undefined, data: Contribution) => {
        if (!goalId) {
            return;
        }
        const createdContribution = await fetchWrapper.post(`${baseUrl}/${goalId}/contribution`, data);
        return setContributions([...contributions, createdContribution.data]);
    }
    
    return {
        getAll,
        create,
    }
}