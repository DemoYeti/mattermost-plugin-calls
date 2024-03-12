SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
        WHERE table_name = 'Calls'
        AND table_schema = DATABASE()
        AND index_name = 'idx_calls_channel_id'
    ) > 0,
    'DROP INDEX idx_calls_channel_id ON Calls;',
    'SELECT 1'
));

PREPARE removeIndexIfExists FROM @preparedStatement;
EXECUTE removeIndexIfExists;
DEALLOCATE PREPARE removeIndexIfExists;

DROP TABLE IF EXISTS Calls;
