package mail

const Confirmation = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Подтверждение уведомлений</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; background-color: #f9f9f9; margin: 0; padding: 20px;">
    <div style="max-width: 600px; margin: 0 auto; background: #ffffff; padding: 20px; border: 1px solid #ddd; border-radius: 8px;">
        <h2 style="color: #555;">Подтверждение уведомлений</h2>
        <p>Перейдите по ссылке в этом письме, если согласны получать уведомления об интересующих Вас мероприятиях платформы от команды <strong>Zero-Cost Developers</strong>:</p>
        <p style="text-align: center; margin: 20px 0;">
            <a href="%s" 
               style="display: inline-block; padding: 12px 20px; color: #fff; background-color: #003b40; text-decoration: none; border-radius: 5px;">Подтвердить</a>
        </p>
        <p style="color: #777;">Если вы не запрашивали это письмо, просто проигнорируйте его.</p>
    </div>
</body>
</html>
`
