<?php

function getClientIP(){
    if (array_key_exists('HTTP_X_FORWARDED_FOR', $_SERVER)){
        return  $_SERVER["HTTP_X_FORWARDED_FOR"];
    }else if (array_key_exists('REMOTE_ADDR', $_SERVER)) {
        return $_SERVER["REMOTE_ADDR"];
    }else if (array_key_exists('HTTP_CLIENT_IP', $_SERVER)) {
        return $_SERVER["HTTP_CLIENT_IP"];
    }

    return '';
}

function goHome() {
	header("Location: /");
	exit();
}

if (count($_POST) == 0) {
	goHome();
}

if (!isset($_POST['email']) || !isset($_POST['course'])) {
	goHome();
}
$email = $_POST['email'];
$tel = isset($_POST['tel']) ? $_POST['tel'] : "nodata";
$course = $_POST['course'];
$name = $_POST['name'];
$ip = getClientIP();
$message = "finalistx.com course\nname: {$name}\nemail: {$email}\nTel: {$tel}\nCourse: {$course}\nIP: <a href='http://ip2loc.com/?ip={$ip}'>{$ip}</a>\n\n";
$email_result = mail('finalist736@gmail.com', 'finalistx.com courses', $message);
if ($email_result) {
	echo '<h2>Email sent!</h2><a href="/">back</a>';
} else {
	echo '<h2>Email error!</h2><a href="/">back</a>';
}

?>

<!DOCTYPE HTML>
<html lang="ru">
<head>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-54235612-1"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag(){dataLayer.push(arguments);}
        gtag('js', new Date());
        gtag('config', 'UA-54235612-1');
    </script>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="Курсы программирования на языке Golang в городе Николаев, Украина.">
    <meta name="keywords" content="курсы, николаев, golang, go, курсы go, курсы golang, программирование, курсы JavaScript, курсы js">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <title>Центр подготовки IT-специалистов Finalist | Николаев</title>
    <!-- Facebook Pixel Code -->
    <script>
        !function(f,b,e,v,n,t,s)
        {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
            n.callMethod.apply(n,arguments):n.queue.push(arguments)};
            if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
            n.queue=[];t=b.createElement(e);t.async=!0;
            t.src=v;s=b.getElementsByTagName(e)[0];
            s.parentNode.insertBefore(t,s)}(window, document,'script',
            'https://connect.facebook.net/en_US/fbevents.js');
        fbq('init', '337075640362725');
        fbq('track', 'PageView');
        fbq('track', 'Purchase');
    </script>
    <noscript><img height="1" width="1" style="display:none"
                   src="https://www.facebook.com/tr?id=337075640362725&ev=PageView&noscript=1"
        /></noscript>
    <!-- End Facebook Pixel Code -->
</head>
<body>
<?php
echo $content;
?>
</body>
</html>