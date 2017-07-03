var app = angular.module("stockTicker", []);

		app.directive('ngEnter', function () {
			return function (scope, element, attrs) {
				element.bind("keydown keypress", function (event) {
					if (event.which === 13) {
						scope.$apply(function () {
							scope.$eval(attrs.ngEnter);
						});

						event.preventDefault();
					}
				});
			};
		});

		app.controller("MainCtl", ["$scope", function ($scope) {
			$scope.messages = [];
			$scope.Ticker = "";
			$scope.Bid = "";
			$scope.Ask = "";
			$scope.Symbol = "";
			$scope.Open = "";
			$scope.Last = "";
			$scope.Date = "";
			var conn = new WebSocket("ws://192.168.33.46/ws");
			// called when the server closes the connection
			conn.onclose = function (e) {
				$scope.$apply(function () {
					$scope.messages.push("DISCONNECTED");
				});
			};
			function IsJsonString(str) {
				try {
					JSON.parse(str);
				} catch (e) {
					return false;
				}
				return true;
			}

			// called when the connection to the server is made
			conn.onopen = function (e) {
				$scope.$apply(function () {
					$scope.messages.push("CONNECTED");
				})
			};

			// called when a message is received from the server
			conn.onmessage = function (e) {
				$scope.$apply(function () {
					var msgReceived = IsJsonString(e.data) ? JSON.parse(e.data) : {
						MessageType: '',
						Message: e.data
					};
					//$scope.messages.push(e.data);
					$scope.Ticker = msgReceived.Ticker;
					$scope.Bid = msgReceived.Bid;
					$scope.Ask = msgReceived.Ask ;
					$scope.Symbol = msgReceived.Symbol;
					$scope.Open = msgReceived.Open;
					$scope.Last = msgReceived.Last;
					$scope.Date = msgReceived.Date;
					console.log(e.data);
				});
			};

			$scope.send = function () {
				conn.send($scope.msg);
				$scope.msg = "";
			}
		}]);