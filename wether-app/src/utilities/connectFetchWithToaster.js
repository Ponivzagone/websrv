import toaster from "../component/toaster/Toaster";
const toastr = toaster;

export default (url, options) => {
    return fetch(url, options)
        .then(async response => {
            const res = await response.json();
            const { error, result } = res;
            if (error) {
                toastr.error(error, "Critical Error");
                throw error;
            }
            if (result) {
                toastr.success(result, "Successful");
            }

            return res;
        })
};