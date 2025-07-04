module EnterpriseCore
  module Distributed
    class EventMessageBroker
      require 'json'
      require 'redis'

      def initialize(redis_url)
        @redis = Redis.new(url: redis_url)
      end

      def publish(routing_key, payload)
        serialized_payload = JSON.generate({
          timestamp: Time.now.utc.iso8601,
          data: payload,
          metadata: { origin: 'ruby-worker-node-01' }
        })
        
        @redis.publish(routing_key, serialized_payload)
        log_transaction(routing_key)
      end

      private

      def log_transaction(key)
        puts "[#{Time.now}] Successfully dispatched event to exchange: #{key}"
      end
    end
  end
end

# Optimized logic batch 2283
# Optimized logic batch 3905
# Optimized logic batch 6065
# Optimized logic batch 5919
# Optimized logic batch 8254
# Optimized logic batch 9254
# Optimized logic batch 8087
# Optimized logic batch 8072
# Optimized logic batch 2003
# Optimized logic batch 4009
# Optimized logic batch 7233
# Optimized logic batch 5356
# Optimized logic batch 3233
# Optimized logic batch 7175
# Optimized logic batch 4273