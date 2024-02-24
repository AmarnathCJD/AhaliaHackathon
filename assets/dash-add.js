$("#add-entry").click(function() {
    Swal.fire({
        title: "Add Entry Details",
        html: `<label for="small-input" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Name</label>
        <input type="text" id="entry-name" class="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"><label for="small-input" class="mt-2 block mb-2 text-sm font-medium text-gray-900 dark:text-white">Quantity</label>
        <input type="text" id="entry-amount" class="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">` +
            '<div class="flex justify-center items-center form-check form-check-inline"><select id="entry-type" class="block py-2.5 px-0 w-auto text-sm text-gray-500 bg-transparent border-0 border-b-2 border-gray-200 appearance-none dark:text-gray-400 dark:border-gray-700 focus:outline-none focus:ring-0 focus:border-gray-200 peer">'+
            
            '<option value="veg">Veg</option>' +
            '<option value="nonveg">NonVeg</option>' +
            '</select></div>',
        showCancelButton: true,
        confirmButtonText: "Add",
        preConfirm: () => {
            return [
                $("#entry-name").val(),
                $("#entry-amount").val(),
                $("#entry-type").val(),
            ];
        }
    }).then(result => {
        if (result.value) {
            const [name, amount, type] = result.value;
            if (name && amount && type) {
                $.post("http://192.168.137.11:8080/api/addentry", {
                    username,
                    name,
                    amount,
                    type,
                }, function(data) {
                    if (data === "success") {
                        Swal.fire("Entry Added", "", "success");
                    } else {
                        Swal.fire("Failed to Add Entry", "", "error");
                    }
                });
            } else {
                Swal.fire("All fields are required", "", "error");
            }
        }
    });
});