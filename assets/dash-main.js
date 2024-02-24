const search_template = `<div class="flex xvz items-center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
<div class="flex items-center space-x-4">
    <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
    <div class="flex flex-col">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">No Results. :(</h3>
        <p class="text-sm text-gray-600 dark:text-gray-400">No suitable food donors found.</p>
    </div>
</div>
</div>
`;

const searchResults = document.getElementById("search-results");

function dispatchSearch() {
  console.log("searching:", $("#search-input").val());
  $.ajax({
    url: "http://192.168.137.11:8080/api/search",
    type: "GET",
    data: {
      name: $("#search-input").val(),
    },
    dataType: "json",
    success: (data) => {
      let html = "";
      if (data == null || data.length == 0 || data == undefined) {
        html = search_template;
        searchResults.innerHTML = html;
        return;
      }
      data.forEach((item) => {
        html += `<div class="flex xvz items-center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
                <a onclick=openEntity('/entity?username=${
                  item.username
                }') class='inline-flex items-center justify-center text-base font-medium text-white rounded-sm hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700 dark:hover:text-white'><div class="flex items-center space-x-4">
                    <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                    <div class="flex flex-col">
                        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">${
                          item.name
                        }</h3>
                        ${
                          item.description
                            ? `<p class="text-sm text-gray-600 dark:text-gray-400">${item.description}</p>`
                            : ""
                        }
                        <p class="text-sm text-gray-600 dark:text-gray-400">
                            ${
                              item.point.rating != 0
                                ? `<i class="fa-solid fa-star text-yellow-400 mr-1"></i>`
                                : ""
                            }
                            <span class="font-semibold">${
                              item.point.rating
                            }</span>
                            <span class="text-gray-400 dark:text-gray-500">•</span>
                            <span class="text-gray-400 dark:text-gray-500">Quantity: ${
                              item.quantity != 0 ? item.quantity : "N/A"
                            }</span>
                        </p>
                    </div>
                </div>
                </a>
            </div>
            `;
      });
      searchResults.innerHTML = html;
    },
    error: () => {
      let html = search_template;
      searchResults.innerHTML = html;
    },
  });
}

// skipcq: JS-0241
$("#search-button").click(dispatchSearch);
$("#search-input").on("keypress", function (e) {
  if (e.key === "Enter") {
    dispatchSearch();
  } else {
    if ($("#search-input").val().length % 3 == 0) {
      dispatchSearch();
    }
  }
});

// skipcq: JS-D1001
function getTopFood() {
  //   skipcq: JS-0125
  $.ajax({
    url: "http://192.168.137.11:8080/api/search",
    type: "GET",
    dataType: "json",
    success: (data) => {
      let html = "";
      let i = data.length > 5 ? 5 : data.length;
      data.forEach((item) => {
        if (item.username == username) return;
        let shouldHide = false;
        i > 0 ? (i -= 1) : (shouldHide = true);
        html += `<div class="flex xvz items-center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition transform scale-50 duration-300 cursor-pointer ${
          shouldHide ? "hidden" : ""
        }"> <a onclick=openEntity('/entity?username=${
          item.username
        }') class='inline-flex items-center justify-center text-base font-medium text-white rounded-sm hover:text-gray-900 hover:bg-gray-00 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-200 dark:hover:text-white hover:shadow-lg dark:hover:shadow-lg hover:border-2 dark:hover:border-2 hover:border-indigo-400 dark:hover:border-gray-200 dark:border-gray-800 border-gray-800'><div class="flex items-center space-x-4">
                        <div class="flex items
                                center space-x-4">
                                <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                                <div class="flex flex-col">
                                        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">${
                                          item.name
                                        }</h3>
                                        ${
                                          item.description
                                            ? `<p class="text-sm text-gray-600 dark:text-gray-400">${item.description}</p>`
                                            : ""
                                        }
                                        <p class="text-sm text-gray-600 dark:text-gray-400">
                                                ${
                                                  //   skipcq: JS-0050
                                                  item.point.rating != 0
                                                    ? // skipcq: JS-R1004
                                                      `<i class="fa-solid fa-star text-yellow-400 mr-1"></i>`
                                                    : ""
                                                }
                                                <span class="font-semibold">${
                                                  item.point.rating
                                                }</span>
                                                <span class="text-gray-400 dark:text-gray-500">•</span>
                                                <span class="text-gray-400 dark:text-gray-500">Quantity: ${
                                                  item.quantity != 0
                                                    ? item.quantity
                                                    : "N/A"
                                                }</span>
                                        </p>
                                </div>
                                </div>
                                </div>
                        </a>
                        </div>
                `;
      });
      searchResults.innerHTML = html;
    },
  });
}

$(document).ready(function () {
  getTopFood();
});

var main_div_backup = document.getElementById("main-div").innerHTML;
function view_orders() {
  main_div_backup = document.getElementById("main-div").innerHTML;
  $.ajax({
    url: "http://192.168.137.11:8080/api/getorders",
    type: "GET",
    dataType: "json",

    success: (data) => {
      html = "";
      data.forEach((item) => {
        var showBtn = true;
        var showBtnDecline = false;
        if (item.approved == false) {
          if (item.orderee == username) {
            console.log("orderee");
            showBtn = true;
          } else if (item.orderer == username) {
            showBtnDecline = true;
            showBtn = false;
          }
        }
        btn_html = "";
        if (!item.approved) {
          if (!showBtnDecline) {
            btn_html += `<button id='btn-${item.id}-app' onclick="approveOrder('${item.id}')" class="px-4 py-2 text-sm font-medium text-white bg-green-500 hover:bg-green-600 dark:bg-green-700 dark:hover:bg-green-800 rounded-lg shadow-md focus:outline-none focus:ring-2 focus:ring-green-500">Request</button>`;
          }
          btn_html += `<button id='btn-${item.id}-dec' onclick="declineOrder('${item.id}')" class="px-4 py-2 text-sm font-medium text-white bg-red-500 hover:bg-red-600 dark:bg-red-700 dark:hover:bg-red-800 rounded-lg shadow-md focus:outline-none focus:ring-2 focus:ring-red-500">Reject</button>`;
        }
        html += `<div class="flex items center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
            <div class="flex items center space-x-4">
                <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                <div class="flex flex-col">
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">${
                      item.product
                    } <b>(x ${item.quantity})</b></h3>
                    <p class="text-sm text-gray-600 dark:text-gray-400">${
                      item.orderer
                    } Registered this Surplus</p>
                    <p class="text-sm text-gray-600 dark:text-gray-400">
                        <i class="fa-solid fa-star text-yellow-400 mr-1"></i>
                        <span class="font-semibold">On ${parseUnixToDateMMDDYY_HHMMSS(
                          item.date_time
                        )}</span>
                        <span class="text-gray-400 dark:text-gray-500">•</span>
                        <span class="text-gray-400 dark:text-gray-500">Status: ${
                          item.approved ? "Sold Out" : "Pending"
                        }</span>
                        <div class="flex items-center space-x-4 mt-2">
                            ${!showBtn ? "" : btn_html}
                        </div>
                    </p>
                </div>
            </div>
        </div>`;
      });
      document.getElementById("main-div").innerHTML = `
  <div class="flex flex-col space-y-4">
    <div class="flex items-center justify-between">
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">Orders</h1>
    </div>
    <div class="flex flex-col space-y-4">
    ${html}
        <div class="flex items
                center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
                <div class="flex items
                        center space-x-4">
                        <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                        <div class="flex flex-col">
                                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Burger</h3>
                                <p class="text-sm text-gray-600 dark:text-gray-400">Burger with extra cheese</p>
                                <p class="text-sm text-gray-600 dark:text-gray-400">
                                        <i class="fa-solid fa-star text-yellow-400 mr-1"></i>
                                        <span class="font-semibold">4.5</span>
                                        <span class="text-gray-400 dark:text-gray-500">•</span>
                                        <span class="text-gray-400 dark:text-gray-500">Quantity: 2</span>
                                </p>
                        </div>
                        </div>
                </div>
        <div class="flex items
                center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
                <div class="flex items
                        center space-x-4">
                        <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                        <div class="flex flex-col">
                                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Pizza</h3>
                                <p class="text-sm text-gray-600 dark:text-gray-400">Pizza with extra cheese</p>
                                <p class="text-sm text-gray-600 dark:text-gray-400">
                                        <i class="fa-solid fa-star text-yellow-400 mr-1"></i>
                                        <span class="font-semibold">4.5</span>
                                        <span class="text-gray-400 dark:text-gray-500">•</span>
                                        <span class="text-gray-400 dark:text-gray-500">Quantity: 1</span>
                                </p>
                        </div>
                        </div>
                </div>
        <div class="flex items

                center justify-between bg-white dark:bg-gray-800 rounded-lg p-4 shadow-lg hover:shadow-xl transition duration-300">
                <div class="flex items
                        center space-x-4">
                        <i class="fa-solid fa-utensils text-3xl text-blue-500 dark:text-yellow-300"></i>
                        <div class="flex flex-col">
                                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Pasta</h3>
                                <p class="text-sm text-gray-600 dark:text-gray-400">Pasta with extra cheese</p>
                                <p class="text-sm text-gray-600 dark:text-gray-400">
                                        <i class="fa-solid fa-star text-yellow-400 mr-1"></i>
                                        <span class="font-semibold">4.5</span>
                                        <span class="text-gray-400 dark:text-gray-500">•</span>
                                        <span class="text-gray-400 dark:text-gray-500">Quantity: 1</span>
                                </p>
                        </div>
                        </div>
                </div>
        </div>
  </div>
  `;

      $("#view-orders").toggleClass("cursor-not-allowed opacity-50", true);
      $("#place-order").toggleClass("cursor-not-allowed opacity-50", false);

      document.getElementById("place-order").disabled = false;
      document.getElementById("view-orders").disabled = true;
    },
  });
}

function approveOrder(id) {
  $.ajax({
    url: "http://192.168.137.11:8080/api/acceptorder",
    type: "GET",
    dataType: "json",
    data: {
      username: username,
      id: id,
    },
    success: (data) => {
      Swal.fire("Order Approved!", "The order has been approved.", "success");
      document.getElementById(`btn-${id}-app`).innerHTML = "Approved!!";
      document.getElementById(`btn-${id}-app`).disabled = true;

      document.getElementById(`btn-${id}-dec`).remove();
    },
    error: (data) => {
      if (data.status == 200) {
        Swal.fire("Order Approved!", "The order has been approved.", "success");
        document.getElementById(`btn-${id}-app`).innerHTML = "Approved!!";
        document.getElementById(`btn-${id}-app`).disabled = true;

        document.getElementById(`btn-${id}-dec`).remove();
      }
    },
  });
}

function declineOrder(id) {
  $.ajax({
    url: "http://192.168.137.11:8080/api/declineorder",
    type: "GET",
    dataType: "json",
    data: {
      id: id,
      username: username,
    },
    success: (data) => {
      document.getElementById(`btn-${id}-dec`).innerHTML = "Declined!!";
      document.getElementById(`btn-${id}-dec`).disabled = true;

      document.getElementById(`btn-${id}-app`).remove();
    },
    error: (data) => {
      if (data.status == 200) {
        document.getElementById(`btn-${id}-dec`).innerHTML = "Declined!!";
        document.getElementById(`btn-${id}-dec`).disabled = true;

        document.getElementById(`btn-${id}-app`).remove();
      }
    },
  });
}

function parseUnixToDateMMDDYY_HHMMSS(unix) {
  var date = new Date(unix * 1000);
  var hours = date.getHours();
  var minutes = "0" + date.getMinutes();
  var seconds = "0" + date.getSeconds();
  var month = date.getMonth() + 1;
  var day = date.getDate();
  var year = date.getFullYear();
  var formattedTime =
    month +
    "/" +
    day +
    "/" +
    year +
    " " +
    hours +
    ":" +
    minutes.substr(-2) +
    ":" +
    seconds.substr(-2);
  return formattedTime;
}

$("#view-orders").click(view_orders);

$("#place-order").click(function () {
  document.getElementById("main-div").innerHTML = main_div_backup;
  $("#search-button").click(dispatchSearch);
  $("#search-input").on("keypress", function (e) {
    if (e.key === "Enter") {
      dispatchSearch();
    } else {
      if ($("#search-input").val().length % 3 == 0) {
        dispatchSearch();
      }
    }
  });

  $("#place-order").toggleClass("cursor-not-allowed opacity-50", true);
  $("#view-orders").toggleClass("cursor-not-allowed opacity-50", false);

  document.getElementById("place-order").disabled = true;
  document.getElementById("view-orders").disabled = false;
});

function openEntity(url) {
  main_div_backup = document.getElementById("main-div").innerHTML;
  orderee = url.split("=")[1];
  $.ajax({
    url: "http://192.168.137.11:8080/api/get",
    type: "GET",
    data: {
      username: url.split("=")[1],
    },
    dataType: "json",
    success: (item) => {
      var totalWaste = "";
      if (item.waste == null) {
        totalWaste = "No Edible Waste Available";
      } else {
        item.waste.forEach((waste) => {
          totalWaste += `<span class="text-gray-400 dark:text-gray-500">${waste.type}: ${waste.quantity}</span>`;
        });
      }
      document.getElementById("main-div").innerHTML = `
      <div>
      <div class="px-4 sm:px-0">
        <h3 class="text-lg font-semibold leading-7 text-gray-900 font-serif">Entity Information</h3>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-gray-600">Business details and application.</p>
      </div>
      <div class="mt-6 border-t border-gray-200">
        <dl class="divide-y divide-gray-200">
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">Name</dt>
            <dd class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">${
              item.name
            }</dd>
          </div>
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">About Entity</dt>
            <dd class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">${
              item.description
            }</dd>
          </div>
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">Entity Category</dt>
            <dd class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">${
              item.e_type
            }</dd>
          </div>
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">GeoLocation</dt>
            <dd class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">${
              item.location
            } (..TODO..)</dd>
          </div>
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">Available Edible Wasted</dt>
            <dd class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">${totalWaste} ${genWasteSlider(
        item.waste
      )}</dd>
          </div>
          <div class="px-4 py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
            <dt class="text-sm font-semibold leading-6 text-gray-800">Select Quantity To Request</dt>
            <dd class="mt-2 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              <button onclick='clickIns()' class="px-6 py-2 mt-0 text-sm font-medium text-white bg-blue-500 hover:bg-blue-600 dark:bg-blue-700 dark:hover:bg-blue-800 rounded-lg shadow-md focus:outline-none focus:ring-2 focus:ring-blue-500">Request For Food</button>
            </dd>
          </div>
        </dl>
      </div>
    </div>
    
`;
    },
  });

  document.getElementById("bat-nan").classList.remove("hidden");
}

function clickIns() {
  Swal.fire({
    title: "Enter Quantity",
    text: "Enter the quantity of food you want to request",
    input: "text",
    showCancelButton: true,
  }).then((result) => {
    if (result.value) {
      // id = random
      id = Math.floor(Math.random() * 1000000);
      $.ajax({
        url: "http://192.168.137.11:8080/api/neworder",
        type: "GET",
        dataType: "json",
        data: {
          id: id,
          orderer: username,
          orderee: orderee,
          product: "Food",
          quantity: result.value,
        },
        success: (data) => {
          Swal.fire("Order Placed!", "Your order has been placed.", "success");
        },
        error: (data) => {
          // if status code is 400
          if (data.status == 400) {
            Swal.fire(
              "Order Failed!",
              "You have already placed an order.",
              "error"
            );
          } else {
            Swal.fire(
              "Order Placed!",
              "Your order has been placed.",
              "success"
            );
          }
        },
      });
    }
  });
}

function goBack() {
  document.getElementById("main-div").innerHTML = main_div_backup;
  document.getElementById("btn-nan").classList.add("hidden");
  $("#search-button").click(dispatchSearch);
  $("#search-input").on("keypress", function (e) {
    if (e.key === "Enter") {
      dispatchSearch();
    } else {
      if ($("#search-input").val().length % 3 == 0) {
        dispatchSearch();
      }
    }
  });
  $("#view-orders").click(view_orders);

  $("#place-order").click(function () {
    document.getElementById("main-div").innerHTML = main_div_backup;
    $("#search-button").click(dispatchSearch);
    $("#search-input").on("keypress", function (e) {
      if (e.key === "Enter") {
        dispatchSearch();
      } else {
        if ($("#search-input").val().length % 3 == 0) {
          dispatchSearch();
        }
      }
    });

    $("#place-order").toggleClass("cursor-not-allowed opacity-50", true);
    $("#view-orders").toggleClass("cursor-not-allowed opacity-50", false);

    document.getElementById("place-order").disabled = true;
    document.getElementById("view-orders").disabled = false;
  });
}

function getIp(latlong) {
  $.ajax({
    url: "http://192.168.137.11:8080/api/ip",
    type: "GET",
    data: {
      ip: latlong,
    },

    success: (data) => {
      console.log(data);
    },
  });
}

function genWasteSlider(waste) {
  if (waste == null) {
    return "- / -";
  }
  total = 0;
  waste.forEach((waste) => {
    total += waste.quantity;
  });
  let html = `<div class="relative mb-6">
  <label for="labels-range-input" class="sr-only">Labels range</label>
  <input id="labels-range-input" type="range" value="1000" min="0" max="${total}" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700">
  <div class="flex justify-between">
    <p class="text-sm text-gray-400 dark:text-gray-500">0</p>
    <p class="text-sm text-gray-400 dark:text-gray-500">${total}</p>
  </div>
  </div>`;

  return html;
}
